#!/bin/bash
set -e

INPUT_ARG="$1"

if [ -z "$INPUT_ARG" ]; then
    echo "Error: No input provided."
    echo "Usage: ./restore.sh <local_file_path> OR <s3_file_key>"
    exit 1
fi

# Check connection vars
if [ -n "$DATABASE_URL" ]; then
    CONNECTION_ARG="$DATABASE_URL"
elif [ -n "$PGDATABASE" ]; then
    CONNECTION_ARG="$PGDATABASE"
else
    echo "Error: Neither DATABASE_URL nor PGDATABASE provided."
    exit 1
fi

# Determine if input is a local file or S3 key
if [ -f "$INPUT_ARG" ]; then
    echo "Local file found: $INPUT_ARG"
    RESTORE_FILE="$INPUT_ARG"
    SHOULD_CLEANUP=false
else
    # Assume S3 key
    if [ -z "$S3_BUCKET" ]; then
        echo "Error: Input file not found locally and S3_BUCKET env var is missing."
        exit 1
    fi
    
    S3_FULL_PATH="s3://${S3_BUCKET}/${S3_PREFIX}${INPUT_ARG}"
    RESTORE_FILE="/tmp/restore_temp.dump"
    SHOULD_CLEANUP=true
    
    echo "Local file '$INPUT_ARG' not found. Attempting download from $S3_FULL_PATH..."
    aws s3 cp "$S3_FULL_PATH" "$RESTORE_FILE"
fi

echo "Restoring data to database..."

# pg_restore flags:
# -a : Data only (no schema)
# -v : Verbose
# -1 : Single transaction
# -d : Connection string/dbname
# --disable-triggers : Disable triggers during data load (speeds up restore, avoids FK issues)
pg_restore -d "$CONNECTION_ARG" -a -v -1 --disable-triggers "$RESTORE_FILE"

if [ "$SHOULD_CLEANUP" = true ]; then
    echo "Cleaning up temp file..."
    rm "$RESTORE_FILE"
fi

echo "Restore complete."
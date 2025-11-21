#!/bin/bash
set -e

# Configuration
TIMESTAMP=$(date +"%Y-%m-%d_%H-%M-%S")
BACKUP_FILE="backup_${TIMESTAMP}.sql"
S3_PATH="s3://${S3_BUCKET}/${S3_PREFIX}${BACKUP_FILE}"

# Determine connection source
if [ -n "$DATABASE_URL" ]; then
    CONNECTION_ARG="$DATABASE_URL"
    echo "Using DATABASE_URL for connection."
elif [ -n "$PGDATABASE" ]; then
    # Fallback to standard libpq env vars (PGHOST, PGUSER, etc.)
    CONNECTION_ARG="$PGDATABASE"
    echo "Using PGDATABASE and standard PG env vars."
else
    echo "Error: Neither DATABASE_URL nor PGDATABASE provided."
    exit 1
fi

if [ -z "$S3_BUCKET" ]; then
    echo "Error: S3_BUCKET is missing."
    exit 1
fi

# Parse TARGET_TABLES (e.g. "users,orders") into flags ("-t users -t orders")
TABLE_FLAGS=""
if [ ! -z "$TARGET_TABLES" ]; then
    for table in $(echo $TARGET_TABLES | tr "," "\n"); do
        TABLE_FLAGS="${TABLE_FLAGS} -t ${table}"
    done
fi

echo "Starting backup to $BACKUP_FILE..."

# pg_dump accepts a URI as the dbname argument
pg_dump "$CONNECTION_ARG" $TABLE_FLAGS -f "/tmp/$BACKUP_FILE"

echo "Backup created. Uploading to $S3_PATH..."

aws s3 cp "/tmp/$BACKUP_FILE" "$S3_PATH"

rm "/tmp/$BACKUP_FILE"

echo "Backup complete."
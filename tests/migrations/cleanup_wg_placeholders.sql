\set ON_ERROR_STOP on
BEGIN;

CREATE TEMP TABLE vehicle_snapshot (
    id text PRIMARY KEY, account_id text, vehicle_id text, type text,
    reference_id text, created_at text, last_battle_time text, battles integer, frame bytea
);
CREATE TEMP TABLE account_snapshot (
    id text PRIMARY KEY, account_id text, type text, reference_id text,
    created_at text, regular_battles integer, regular_frame bytea,
    rating_battles integer DEFAULT 20, rating_frame bytea DEFAULT '\x1234'
);
CREATE FUNCTION pg_temp.stats(battles integer, wins integer, damage bigint) RETURNS bytea
LANGUAGE sql AS $$
    SELECT convert_to(jsonb_build_object('battles', battles::text,
        'battlesWon', wins::text, 'damageDealt', damage::text, 'mmRating', '3000')::text, 'UTF8');
$$;

INSERT INTO vehicle_snapshot VALUES
('bad1', 'a', '7953', 'daily', 'shared', '2026-09-01T00:00:00Z', '1970-01-01T00:00:00Z', 98, pg_temp.stats(98, 61, 0)),
('bad2', 'a', '7953', 'daily', 'shared', '2026-09-02T00:00:00Z', '1970-01-01T00:00:00Z', 99, pg_temp.stats(99, 62, 0)),
('repaired', 'a', '7953', 'daily', 'shared', '2026-09-03T00:00:00Z', '2026-09-02T12:00:00Z', 100, pg_temp.stats(100, 63, 240525)),
('other-account', 'b', '7953', 'daily', 'shared', '2026-09-01T01:00:00Z', '2026-09-01T00:00:00Z', 300, pg_temp.stats(300, 200, 500000)),
('other-reference', 'a', '7953', 'daily', 'other', '2026-09-01T01:00:00Z', '2026-09-01T00:00:00Z', 300, pg_temp.stats(300, 200, 500000)),
('other-type', 'a', '7953', 'widget', 'shared', '2026-09-01T01:00:00Z', '2026-09-01T00:00:00Z', 300, pg_temp.stats(300, 200, 500000)),
('zero-damage', 'a', '1', 'daily', 'shared', '2026-09-01T00:00:00Z', '2026-08-31T12:00:00Z', 1, pg_temp.stats(1, 0, 0)),
('unplayed', 'a', '2', 'daily', 'shared', '2026-09-01T00:00:00Z', '1970-01-01T00:00:00Z', 0, pg_temp.stats(0, 0, 0)),
('unexpected-details', 'a', '3', 'daily', 'shared', '2026-09-01T00:00:00Z', '1970-01-01T00:00:00Z', 1, pg_temp.stats(1, 0, 100));

INSERT INTO account_snapshot (id, account_id, type, reference_id, created_at, regular_battles, regular_frame) VALUES
('before', 'a', 'daily', 'shared', '2026-08-31T00:00:00Z', 10, pg_temp.stats(10, 5, 10000)),
('during1', 'a', 'daily', 'shared', '2026-09-01T00:00:00Z', 108, pg_temp.stats(108, 66, 10000)),
('during2', 'a', 'daily', 'shared', '2026-09-02T00:00:00Z', 111, pg_temp.stats(111, 68, 14000)),
('after', 'a', 'daily', 'shared', '2026-09-03T00:00:00Z', 112, pg_temp.stats(112, 69, 254525)),
('other-account', 'b', 'daily', 'shared', '2026-09-02T00:00:00Z', 300, pg_temp.stats(300, 200, 500000)),
('other-reference', 'a', 'daily', 'other', '2026-09-02T00:00:00Z', 300, pg_temp.stats(300, 200, 500000)),
('other-type', 'a', 'widget', 'shared', '2026-09-02T00:00:00Z', 300, pg_temp.stats(300, 200, 500000));

CREATE TEMP TABLE expected_accounts AS SELECT * FROM account_snapshot;
UPDATE expected_accounts SET regular_battles = 10, regular_frame = pg_temp.stats(10, 5, 10000) WHERE id = 'during1';
UPDATE expected_accounts SET regular_battles = 12, regular_frame = pg_temp.stats(12, 6, 14000) WHERE id = 'during2';
CREATE TEMP TABLE expected_vehicles AS SELECT * FROM vehicle_snapshot WHERE id NOT IN ('bad1', 'bad2');

\ir ../../internal/database/migrations/20260926190000.up.sql
\ir ../../internal/database/migrations/20260926190000.up.sql

DO $$ BEGIN
    IF EXISTS (
        (SELECT * FROM account_snapshot EXCEPT SELECT * FROM expected_accounts)
        UNION ALL
        (SELECT * FROM expected_accounts EXCEPT SELECT * FROM account_snapshot)
    ) THEN RAISE EXCEPTION 'Account totals or rating data changed incorrectly'; END IF;
    IF EXISTS (
        (SELECT * FROM vehicle_snapshot EXCEPT SELECT * FROM expected_vehicles)
        UNION ALL
        (SELECT * FROM expected_vehicles EXCEPT SELECT * FROM vehicle_snapshot)
    ) THEN RAISE EXCEPTION 'Unexpected vehicle snapshots removed'; END IF;
END $$;
ROLLBACK;

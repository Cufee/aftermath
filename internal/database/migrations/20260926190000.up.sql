-- These WG placeholders contain battles/wins, but no detailed statistics.
WITH invalid AS MATERIALIZED (
    SELECT id, account_id, vehicle_id, type, reference_id, created_at, battles,
           (convert_from(frame, 'UTF8')::jsonb ->> 'battlesWon')::bigint AS wins
    FROM vehicle_snapshot
    WHERE battles > 0
      AND last_battle_time::timestamptz <= 'epoch'::timestamptz
      AND NOT EXISTS (
          SELECT 1
          FROM jsonb_each_text(convert_from(frame, 'UTF8')::jsonb
                               - ARRAY['battles', 'battlesWon', 'mmRating']) AS counter
          WHERE counter.value::numeric <> 0
      )
), corrections AS (
    SELECT a.id, sum(v.battles)::bigint AS battles, sum(v.wins)::bigint AS wins
    FROM account_snapshot a
    JOIN invalid v ON v.account_id = a.account_id
                  AND v.type = a.type
                  AND v.reference_id = a.reference_id
                  AND v.created_at <= a.created_at
    -- Stop correcting a tank once a newer snapshot supplies its statistics.
    WHERE NOT EXISTS (
        SELECT 1 FROM vehicle_snapshot newer
        WHERE newer.account_id = v.account_id
          AND newer.vehicle_id = v.vehicle_id
          AND newer.type = v.type
          AND newer.reference_id = v.reference_id
          AND newer.created_at > v.created_at
          AND newer.created_at <= a.created_at
    )
    GROUP BY a.id
), corrected AS (
    UPDATE account_snapshot a
    SET regular_battles = a.regular_battles - c.battles,
        regular_frame = convert_to((
            convert_from(a.regular_frame, 'UTF8')::jsonb || jsonb_build_object(
                'battles', ((convert_from(a.regular_frame, 'UTF8')::jsonb ->> 'battles')::bigint - c.battles)::text,
                'battlesWon', ((convert_from(a.regular_frame, 'UTF8')::jsonb ->> 'battlesWon')::bigint - c.wins)::text
            )
        )::text, 'UTF8')
    FROM corrections c
    WHERE a.id = c.id
    RETURNING a.id
)
DELETE FROM vehicle_snapshot v USING invalid i WHERE v.id = i.id;

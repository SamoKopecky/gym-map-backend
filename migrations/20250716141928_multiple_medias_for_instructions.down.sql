ALTER TABLE instruction
RENAME COLUMN media_ids TO media_id;

ALTER TABLE instruction
ALTER COLUMN media_id DROP NOT NULL;

ALTER TABLE instruction
ALTER COLUMN media_id DROP DEFAULT;

ALTER TABLE instruction
ALTER COLUMN media_id TYPE integer
USING CASE
    WHEN media_id IS NULL THEN NULL                       -- Handles if original array was NULL (unlikely if 'up' made it NOT NULL and default '{}')
    WHEN array_length(media_id, 1) = 0 THEN NULL          -- Handles empty arrays (e.g., {})
    ELSE media_id[1]                                      -- Takes the first element
END;

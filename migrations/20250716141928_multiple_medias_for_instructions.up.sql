ALTER TABLE instruction
ALTER COLUMN media_id TYPE integer[] 
USING CASE
    WHEN media_id IS NULL THEN NULL
    ELSE ARRAY[media_id] -- Convert the single integer into an array containing that integer
END;

ALTER TABLE instruction
ALTER COLUMN media_id SET DEFAULT '{}';

ALTER TABLE instruction
ALTER COLUMN media_id SET NOT NULL;

ALTER TABLE instruction
RENAME COLUMN media_id TO media_ids;

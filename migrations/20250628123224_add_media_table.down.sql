DROP TABLE IF EXISTS media;

ALTER TABLE instruction
ADD COLUMN file_id VARCHAR(36);

ALTER TABLE instruction
ADD COLUMN file_name TEXT;

ALTER TABLE instruction
DROP COLUMN media_id;

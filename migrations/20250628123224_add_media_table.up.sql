CREATE TABLE IF NOT EXISTS media (
  id SERIAL PRIMARY KEY,
  original_file_name TEXT NOT NULL,
  disk_file_name VARCHAR(255) NOT NULL,
  content_type VARCHAR(64) NOT NULL,
  created_at TIMESTAMP
  WITH
    TIME ZONE NOT NULL,
    updated_at TIMESTAMP
  WITH
    TIME ZONE NOT NULL
);

ALTER TABLE instruction
DROP COLUMN file_id;

ALTER TABLE instruction
DROP COLUMN file_name;

ALTER TABLE instruction
ADD COLUMN media_id INTEGER;

ALTER TABLE media
RENAME COLUMN original_file_name TO name;

ALTER TABLE media
RENAME COLUMN disk_file_name TO path;

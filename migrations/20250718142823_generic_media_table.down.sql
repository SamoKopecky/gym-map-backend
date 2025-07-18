ALTER TABLE media
RENAME COLUMN name TO original_file_name;

ALTER TABLE media
RENAME COLUMN path TO disk_file_name;

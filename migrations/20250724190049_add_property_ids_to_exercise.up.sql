ALTER TABLE exercise
ADD COLUMN property_ids integer[] NOT NULL DEFAULT '{}';

ALTER TABLE exercise
DROP COLUMN muscle_groups;

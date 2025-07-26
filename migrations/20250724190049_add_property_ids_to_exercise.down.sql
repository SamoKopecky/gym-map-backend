ALTER TABLE exercise
DROP COLUMN property_ids;

ALTER TABLE exercise
ADD COLUMN muscle_groups TEXT[] DEFAULT '{}';

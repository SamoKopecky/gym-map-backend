ALTER TABLE exercise
DROP CONSTRAINT fk_exercise_machines;

ALTER TABLE instruction
DROP CONSTRAINT fk_instruction_exercise;

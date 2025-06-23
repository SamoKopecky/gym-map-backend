ALTER TABLE exercise ADD CONSTRAINT fk_exercise_machines FOREIGN KEY (machine_id) REFERENCES machine (id) ON DELETE CASCADE;

ALTER TABLE instruction ADD CONSTRAINT fk_instruction_exercise FOREIGN KEY (exercise_id) REFERENCES exercise (id) ON DELETE CASCADE;

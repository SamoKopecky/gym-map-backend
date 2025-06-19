CREATE TABLE IF NOT EXISTS instruction (
  id SERIAL PRIMARY KEY,
  exercise_id INTEGER NOT NULL,
  user_id VARCHAR(36) NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP
  WITH
    TIME ZONE NOT NULL,
    updated_at TIMESTAMP
  WITH
    TIME ZONE NOT NULL
);

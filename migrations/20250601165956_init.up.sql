CREATE TABLE IF NOT EXISTS machine (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    muscle_groups TEXT[],
    width INTEGER NOT NULL,
    height INTEGER NOT NULL,
    position_x INTEGER NOT NULL,
    position_y INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL 
);

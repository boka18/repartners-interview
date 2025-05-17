CREATE TABLE IF NOT EXISTS pack_sizes (
    id SERIAL PRIMARY KEY,
    size INT NOT NULL UNIQUE
);

-- Initial data
INSERT INTO pack_sizes (size) VALUES (23), (31), (53)
    ON CONFLICT (size) DO NOTHING;

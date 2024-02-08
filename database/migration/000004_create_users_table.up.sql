CREATE TABLE IF NOT EXISTS Users
(
    id SERIAL PRIMARY KEY,
    email TEXT,
    name VARCHAR(100),
    phone NUMERIC(10),
    address VARCHAR(150)
)
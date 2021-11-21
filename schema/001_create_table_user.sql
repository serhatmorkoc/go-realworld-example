CREATE TABLE IF NOT EXISTS users(
    user_id SERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    bio TEXT,
    image VARCHAR,
    created_at Timestamp NOT NULL,
    updated_at Timestamp NOT NULL,
    UNIQUE(username),
    UNIQUE(email),
);
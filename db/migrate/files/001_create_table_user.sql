CREATE SEQUENCE IF NOT EXISTS users_seq;

CREATE TABLE IF NOT EXISTS users(
    user_id int DEFAULT NEXTVAL ('users_seq') NOT NULL,
    email VARCHAR NOT NULL,
    token VARCHAR NOT NULL,
    username VARCHAR NOT NULL,
    bio TEXT,
    image VARCHAR,
    created_at Timestamp NOT NULL,
    updated_at Timestamp NOT NULL,
    CONSTRAINT PK_users PRIMARY KEY(user_id)
    );
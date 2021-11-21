CREATE TABLE IF NOT EXISTS comments(
    comment_id SERIAL PRIMARY KEY,
    article_id INTEGER NOT NULL,
    body TEXT NOT NULL,
    author VARCHAR NOT NULL,
    created_at Timestamp NOT NULL,
    updated_at Timestamp NOT NULL,
);
CREATE TABLE IF NOT EXISTS articles(
    article_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    slug VARCHAR NOT NULL,
    title VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    body TEXT NOT NULL,
    tag_list VARCHAR NOT NULL,
    favorited BOOLEAN NOT NULL,
    favorites_count INTEGER NOT NULL,
    created_at Timestamp NOT NULL,
    updated_at Timestamp NOT NULL,
);
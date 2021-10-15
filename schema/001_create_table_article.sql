CREATE SEQUENCE IF NOT EXISTS articles_seq;

CREATE TABLE IF NOT EXISTS articles
(
    article_id int DEFAULT NEXTVAL ('articles_seq') NOT NULL,
    slug VARCHAR NOT NULL,
    title VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    body TEXT NOT NULL,
    tag_list VARCHAR NOT NULL,
    favorited BOOLEAN NOT NULL,
    favorites_count int NOT NULL,
    created_at Timestamp NOT NULL,
    updated_at Timestamp NOT NULL,
    CONSTRAINT PK_articles PRIMARY KEY
    (article_id)
);
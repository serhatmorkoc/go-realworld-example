CREATE SEQUENCE IF NOT EXISTS comments_seq;

CREATE TABLE IF NOT EXISTS comments
(
    comment_id int DEFAULT NEXTVAL ('comments_seq') NOT NULL,
    article_id int NOT NULL,
    body TEXT NOT NULL,
    author VARCHAR NOT NULL,
    created_at Timestamp NOT NULL,
    updated_at Timestamp NOT NULL,
    CONSTRAINT PK_comments PRIMARY KEY
    (comment_id)
);
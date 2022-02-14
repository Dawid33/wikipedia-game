CREATE SCHEMA guestbook;
CREATE TABLE guestbook.posts (
    post_id    SERIAL UNIQUE PRIMARY KEY,
    name        text,
    comment     text,
    post_time   timestamp
);
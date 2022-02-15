CREATE SCHEMA game;
CREATE TABLE game.sessions (
    session_id  uuid default uuid_generate_v4(),
    post_time   timestamp default CURRENT_TIMESTAMP

);
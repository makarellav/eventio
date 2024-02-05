CREATE TABLE events
(
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT NOT NULL,
    location    TEXT NOT NULL,
    date_time   DATETIME NOT NULL,
    user_id     INT,
);
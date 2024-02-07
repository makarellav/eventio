CREATE TABLE events
(
    id          INT PRIMARY KEY AUTOINCREMENT,
    name        TEXT     NOT NULL,
    description TEXT     NOT NULL,
    location    TEXT     NOT NULL,
    date_time   DATETIME NOT NULL,
    user_id     INT
);
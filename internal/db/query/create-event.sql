-- name: CreateEvent :one
INSERT INTO events(name, description, location, date_time, user_id)
VALUES (?, ?, ?, ?, ?)
RETURNING *;
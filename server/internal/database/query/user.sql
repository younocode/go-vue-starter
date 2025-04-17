-- name: GetUserByEmail :one
SELECT id, password, email
FROM users
WHERE email = $1;

-- name: IsEmailAvailable :one
SELECT NOT EXISTS (SELECT 1
                   from users
                   WHERE email = $1);

-- name: CreateUser :one
INSERT INTO users (email,
                   password)
VALUES ($1, $2)
RETURNING *;

-- name: UpdatePasswordByEmail :one
UPDATE users
SET password = $1,
    updated_at    = $2
WHERE email = $3
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUser :exec
UPDATE users
set email = $2,
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
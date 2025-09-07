-- name: CreateUser :one
INSERT INTO recycle.user (id, email, hashed_password, confirmation_token, role, created_by, updated_by)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: ConfirmUserAccess :exec
UPDATE recycle.user SET confirmed_at = $2, updated_at = now(), updated_by = $3 WHERE id = $1;

-- name: FindUserById :one
SELECT id, email, hashed_password, confirmed_at, confirmation_token, role, created_at, updated_at, deleted_at, created_by, updated_by
FROM recycle.user WHERE id = $1;

-- name: FindUserByEmail :one
SELECT id, email, hashed_password, confirmed_at, confirmation_token, role, created_at, updated_at, deleted_at, created_by, updated_by
FROM recycle.user WHERE email = $1;

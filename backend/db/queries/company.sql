-- name: CreateCompany :one
INSERT INTO recycle.company (id, name, company_owner_id, created_by, updated_by)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: FindCompanyById :one
SELECT
    sqlc.embed(company),
    sqlc.embed(users)
FROM recycle.company company
INNER JOIN recycle.user users ON company.company_owner_id = users.id
WHERE company.id = $1;

-- name: GetAdmin :one
SELECT * FROM admin
WHERE id = $1 LIMIT 1;

-- name: GetAdminByEmail :one
SELECT * FROM admin
WHERE email = $1 LIMIT 1;

-- name: ListAdmin :many
SELECT 
    id,
    email,
    first_name,
    last_name,
    permission,
    last_accessed_at,
    created_at,
    deleted_at 
FROM admin
ORDER BY first_name;
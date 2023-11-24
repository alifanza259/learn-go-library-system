-- name: GetMember :one
SELECT * FROM members
WHERE id = $1 LIMIT 1;

-- name: ListMembers :many
SELECT 
  id,
  email,
  first_name,
  last_name,
  dob,
  gender,
  created_at,
  updated_at,
  deleted_at
FROM members
ORDER BY first_name;

-- name: CreateMember :one
INSERT INTO members (
  email, 
  first_name, 
  last_name, 
  dob,
  gender,
  password
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetMemberByEmail :one
SELECT * FROM members
WHERE email = $1 LIMIT 1;

-- name: UpdateMember :one
UPDATE members SET 
  email_verified_at=coalesce(sqlc.narg('email_verified_at'), email_verified_at)
WHERE id=sqlc.arg(id)
RETURNING *;
-- name: GetMember :one
SELECT * FROM members
WHERE id = $1 LIMIT 1;

-- name: ListMembers :many
SELECT * FROM members
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
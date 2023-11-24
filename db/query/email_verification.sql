-- name: CreateEmailVerification :one
INSERT INTO email_verifications (
  member_id, 
  token
) VALUES (
  $1, $2
)
RETURNING *;
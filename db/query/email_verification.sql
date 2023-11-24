-- name: CreateEmailVerification :one
INSERT INTO email_verifications (
  member_id, 
  token
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEmailVerification :one
SELECT * FROM email_verifications 
WHERE token = $1 LIMIT 1;

-- name: UpdateEmailVerification :exec
UPDATE email_verifications SET
  is_used = $2
WHERE token = $1;
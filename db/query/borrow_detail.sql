-- name: GetBorrow :one
SELECT * FROM borrow_details
WHERE id = $1 LIMIT 1;

-- name: CreateBorrow :one
INSERT INTO borrow_details (
  book_id,
  borrowed_at,
  returned_at
) VALUES (
  $1, $2, $3
)
RETURNING *;
-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: GetTransactionAndBorrowDetail :one
SELECT t.id trx_id, t.member_id trx_member_id, bd.id bd_id FROM transactions t
JOIN borrow_details bd ON t.borrow_id = bd.id
WHERE t.id = $1 LIMIT 1;

-- name: GetTransactionAssociatedDetail :one
SELECT t.id trx_id, b.title b_title, m.email, m.first_name, b.id b_id, t.purpose t_purpose FROM transactions t
JOIN borrow_details bd ON t.borrow_id = bd.id
JOIN books b ON bd.book_id = b.id
JOIN members m ON t.member_id = m.id
WHERE t.id = $1 LIMIT 1;

-- name: CreateTransaction :one
INSERT INTO transactions (
  member_id,
  borrow_id,
  purpose,
  status
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateTransaction :one
UPDATE transactions 
SET 
  admin_id=$1,
  status=$2,
  note=$3
WHERE id=$4
RETURNING *;

-- name: GetBorrowHistory :many
SELECT t.ID t_id, b.title b_title, b.author b_author, b.image_url b_image_url, bd.borrowed_at bd_borrowed_at, bd.returned_at bd_returned_at FROM transactions t
JOIN borrow_details bd ON t.borrow_id = bd.id
JOIN books b ON bd.book_id = b.id
WHERE t.member_id = sqlc.arg('member_id') AND t.status = coalesce(sqlc.narg('status'), status) ;
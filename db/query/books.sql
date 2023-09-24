-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT *
FROM books
ORDER BY title;

-- name: CreateBook :one
INSERT INTO books (
  isbn, 
  title, 
  description, 
  author,
  image_url,
  genre,
  quantity,
  published_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;
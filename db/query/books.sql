-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: GetBookForUpdate :one
SELECT * FROM books
WHERE id = $1 LIMIT 1 FOR UPDATE;

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

-- name: UpdateBook :one
UPDATE books 
SET 
  isbn=coalesce(sqlc.narg('isbn'), isbn),
  title=coalesce(sqlc.narg('title'), title),
  description=coalesce(sqlc.narg('description'), description),
  author=coalesce(sqlc.narg('author'), author),
  image_url=coalesce(sqlc.narg('image_url'), image_url),
  genre=coalesce(sqlc.narg('genre'), genre),
  quantity=coalesce(sqlc.narg('quantity'), quantity),
  published_at=coalesce(sqlc.narg('published_at'), published_at)
WHERE id=sqlc.arg(id)
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id=$1;
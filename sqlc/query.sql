-- name: GetBook :one
SELECT * FROM books
WHERE isbn = ? LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY author;

-- name: InsertBook :exec
INSERT INTO books (isbn, title, description, author, image_url)
VALUES (?, ?, ?, ?, ?);

-- name: InsertManyBooks :exec
INSERT INTO books (isbn, title, description, author, image_url)
VALUES
    (?, ?, ?, ?, ?),
    (?, ?, ?, ?, ?),
    (?, ?, ?, ?, ?),
    (?, ?, ?, ?, ?),
    (?, ?, ?, ?, ?);

-- name: UpdateBook :exec
UPDATE books
SET 
    title = COALESCE(?, title),
    description = COALESCE(?, description),
    author = COALESCE(?, author),
    image_url = COALESCE(?, image_url),
    isbn = COALESCE(?, isbn)
WHERE
    id = ?;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id= ?;

-- name: CountBooks :one
SELECT COUNT(*) FROM books;

-- name: GetBooksByAuthor :many
SELECT * FROM books
WHERE author = ?;

-- name: GetBookByID :one
SELECT * FROM books
WHERE id = ?;
-- name: CreateCategory :one
INSERT INTO categories (
    user_id,
    title,
    type,
    description
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories 
WHERE id = $1 LIMIT 1;

-- name: GetCategories :many
SELECT * FROM categories
WHERE
  user_id = $1
AND
  type = $2
AND
  title LIKE CONCAT('%', @title::text, '%')
AND
  description LIKE CONCAT('%', @description::text, '%');


-- name: GetCategoriesByUserIdAndType :many
SELECT * FROM categories 
WHERE user_id = $1 AND type = $2; 

-- name: GetCategoriesByUserIdAndTypeAndTitle :many
SELECT * FROM categories 
WHERE user_id = $1 AND type = $2 
AND title LIKE $3;

-- name: GetCategoriesByUserIdAndTypeAndDescription :many
SELECT * FROM categories 
WHERE user_id = $1 AND type = $2 
AND description LIKE $3;

-- name: UpdateCategories :one
UPDATE categories 
SET title = $2, description = $3 
WHERE id = $1 
RETURNING *;

-- name: DeleteCategories :exec
DELETE FROM categories 
WHERE id = $1;
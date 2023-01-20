-- name: CreateAccount :one
INSERT INTO accounts (
    user_id,
    category_id,
    title,
    type,
    description,
    value,
    date
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts 
WHERE id = $1 LIMIT 1;

-- name: GetAccounts :many
SELECT 
    a.id,
    a.user_id,
    a.title,
    a.type,
    a.description,
    a.value,
    a.date,
    a.created_at,
    c.title as category_title
FROM accounts a 
LEFT JOIN categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2 
AND a.category_id = $3 AND a.title LIKE $4 
AND a.description LIKE $5 AND a.date = $6;

-- name: GetAccountsByUserIdAndType :many
SELECT 
    a.id,
    a.user_id,
    a.title,
    a.type,
    a.description,
    a.value,
    a.date,
    a.created_at,
    c.title as category_title
FROM accounts a 
LEFT JOIN categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2; 

-- name: GetAccountsByUserIdAndTypeAndCategoryId :many
SELECT 
    a.id,
    a.user_id,
    a.title,
    a.type,
    a.description,
    a.value,
    a.date,
    a.created_at,
    c.title as category_title
FROM accounts a 
LEFT JOIN categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2 
AND a.category_id = $3;

-- name: GetAccountsByUserIdAndTypeAndCategoryIdAndTitle :many
SELECT 
    a.id,
    a.user_id,
    a.title,
    a.type,
    a.description,
    a.value,
    a.date,
    a.created_at,
    c.title as category_title
FROM accounts a 
LEFT JOIN categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2 
AND a.category_id = $3 AND a.title LIKE $4;

-- name: GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescription :many
SELECT 
    a.id,
    a.user_id,
    a.title,
    a.type,
    a.description,
    a.value,
    a.date,
    a.created_at,
    c.title as category_title
FROM accounts a 
LEFT JOIN categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2 
AND a.category_id = $3 AND a.title LIKE $4
AND a.description LIKE $5;

-- name: GetAccountsByUserIdAndTypeAndTitle :many
SELECT 
    a.id,
    a.user_id,
    a.title,
    a.type,
    a.description,
    a.value,
    a.date,
    a.created_at,
    c.title as category_title
FROM accounts a 
LEFT JOIN categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2
AND a.title LIKE $3; 

-- name: GetAccountsByUserIdAndTypeAndDescription :many
SELECT 
    a.id,
    a.user_id,
    a.title,
    a.type,
    a.description,
    a.value,
    a.date,
    a.created_at,
    c.title as category_title
FROM accounts a 
LEFT JOIN categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2
AND a.description LIKE $3; 

-- name: GetAccountsByUserIdAndTypeAndDate :many
SELECT 
    a.id,
    a.user_id,
    a.title,
    a.type,
    a.description,
    a.value,
    a.date,
    a.created_at,
    c.title as category_title
FROM accounts a 
LEFT JOIN categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2
AND a.date LIKE $3; 

-- name: GetAccountsReports :one
SELECT SUM(value) AS sum_value FROM accounts
WHERE user_id = $1 AND type = $2;

-- name: GetAccountsGraph :one
SELECT COUNT(*) FROM accounts
WHERE user_id = $1 AND type = $2;

-- name: UpdateAccount :one
UPDATE accounts 
SET title = $2, description = $3, value = $4
WHERE id = $1 
RETURNING *;

-- name: DeleteAccounty :exec
DELETE FROM accounts 
WHERE id = $1;
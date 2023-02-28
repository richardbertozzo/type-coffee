-- name: GetCoffeeById :one
SELECT * FROM coffee
WHERE id = $1;
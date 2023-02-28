// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: coffee.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const getCoffeeById = `-- name: GetCoffeeById :one
SELECT id, specie, owner, country_of_origin, company, aroma, flavor, aftertaste, acidity, body, sweetness FROM coffee
WHERE id = $1
`

func (q *Queries) GetCoffeeById(ctx context.Context, id uuid.UUID) (Coffee, error) {
	row := q.db.QueryRow(ctx, getCoffeeById, id)
	var i Coffee
	err := row.Scan(
		&i.ID,
		&i.Specie,
		&i.Owner,
		&i.CountryOfOrigin,
		&i.Company,
		&i.Aroma,
		&i.Flavor,
		&i.Aftertaste,
		&i.Acidity,
		&i.Body,
		&i.Sweetness,
	)
	return i, err
}
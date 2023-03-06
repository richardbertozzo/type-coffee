-- name: GetCoffeeById :one
SELECT * FROM coffee
WHERE id = $1;

-- name: InsertCoffee :one
INSERT INTO coffee (
    specie,
    owner,
    country_of_origin,
    company,
    aroma,
    flavor,
    aftertaste,
    acidity,
    body,
    sweetness
) VALUES (
    @specie::varchar,
    @owner::varchar,
    @country_of_origin::varchar,
    @company::varchar,
    @aroma,
    @flavor,
    @aftertaste,
    @acidity,
    @body,
    @sweetness
) RETURNING id;
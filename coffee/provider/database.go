package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/richardbertozzo/type-coffee/coffee"
)

type databaseService struct {
	db DBTX
}

func NewDatabase(db DBTX) coffee.Service {
	return &databaseService{
		db: db,
	}
}

func (d *databaseService) GetCoffeeOptionsByCharacteristics(ctx context.Context, filter coffee.Filter) ([]coffee.OptionProvider, error) {
	orderByArgs := make([]string, len(filter.Characteristics))
	for i, characteristic := range filter.Characteristics {
		orderByArgs[i] = fmt.Sprintf("%s %s", string(characteristic), getSort(filter.Sort))
	}
	orderBy := strings.Join(orderByArgs, ",")

	query := `
		SELECT * FROM coffee
		ORDER BY $1
		LIMIT $2
	`
	rows, err := d.db.Query(ctx, query, orderBy, filter.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var coffees []Coffee
	for rows.Next() {
		var c Coffee
		if err := rows.Scan(
			&c.ID,
			&c.Specie,
			&c.Owner,
			&c.CountryOfOrigin,
			&c.Company,
			&c.Aroma,
			&c.Flavor,
			&c.Aftertaste,
			&c.Acidity,
			&c.Body,
			&c.Sweetness,
		); err != nil {
			return nil, err
		}
		coffees = append(coffees, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	options := make([]coffee.OptionProvider, len(coffees))
	for i, c := range coffees {
		options[i] = coffee.OptionProvider{
			Message: buildMessageCoffee(c),
		}
	}

	return options, nil
}

func getSort(sort bool) string {
	if sort {
		return "ASC"
	}
	return "DESC"
}

func buildMessageCoffee(coffee Coffee) string {
	// TODO: improve this message
	return fmt.Sprintf("you might taste a coffee from %s - specie %s and from the owner %s", coffee.CountryOfOrigin, coffee.Specie, coffee.Owner)
}

package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgtype/pgxtype"

	"github.com/richardbertozzo/type-coffee/coffee"
)

type databaseService struct {
	queries *Queries
}

func NewDatabaseService(querier pgxtype.Querier) coffee.Service {
	return &databaseService{
		queries: New(querier),
	}
}

func (d *databaseService) GetCoffeeOptionsByCharacteristics(ctx context.Context, filter coffee.Filter) ([]coffee.OptionProvider, error) {
	// TODO: implement me
	_, err := d.queries.GetCoffeeById(ctx, uuid.New())
	return nil, err
}

package usecase

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"

	"github.com/richardbertozzo/type-coffee/coffee"
)

const disclaimerBestCoffeeMsg = `There's no "best" coffee, as it's entirely subjective!  What someone finds delicious, another might find awful. 

Consider your preferences and explore different coffee types

Here's how to find your perfect flavor:

* **Start with a sample pack:** Many online retailers offer small bags of various coffees.
* **Visit a local coffee shop:**  They often have a wide selection and can give recommendations.
* **Experiment:**  Don't be afraid to try different roasts and origins until you find what you love.

Remember, the "best" is the one YOU enjoy the most! Happy coffee tasting! 
`

type useCase struct {
	geminiAPIProvider coffee.Service
	dbService         coffee.Service
	dbServiceEnabled  bool
}

// New returns a new coffee use case
func New(geminiProvider, dbService coffee.Service) *useCase {
	return &useCase{
		geminiAPIProvider: geminiProvider,
		dbService:         dbService,
		dbServiceEnabled:  dbService != nil,
	}
}

func (u useCase) GetBestCoffees(ctx context.Context, filter coffee.Filter) (*coffee.BestCoffees, error) {
	lenChar := len(filter.Characteristics)
	if lenChar == 0 {
		return nil, errors.New("must have minimum 1 input characteristic")
	} else if lenChar > 3 {
		return nil, errors.New("must have up to 3 input characteristics")
	}

	filter.Limit = setDefaultLimit(filter.Limit)

	g, ctx := errgroup.WithContext(ctx)

	var geminiOptions []coffee.Option
	g.Go(func() error {
		opts, err := u.geminiAPIProvider.GetCoffeeOptionsByCharacteristics(ctx, filter)
		if err != nil {
			return err
		}

		for _, opt := range opts {
			geminiOptions = append(geminiOptions, coffee.Option{
				Message: opt.Message,
			})
		}

		return nil
	})

	var dbOpts []coffee.Option
	if u.dbServiceEnabled {
		g.Go(func() error {
			opts, err := u.dbService.GetCoffeeOptionsByCharacteristics(ctx, filter)
			if err != nil {
				return err
			}

			for _, opt := range opts {
				dbOpts = append(dbOpts, coffee.Option{
					Message: opt.Message,
				})
			}

			return nil
		})

	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &coffee.BestCoffees{
		Characteristics: filter.Characteristics,
		Gemini:          &geminiOptions,
		Database:        &dbOpts,
		Disclaimer:      disclaimerBestCoffeeMsg,
	}, nil
}

func setDefaultLimit(limit int) int {
	defaultLimit := 5
	if limit == 0 || limit > 5 {
		return defaultLimit
	}
	return limit
}

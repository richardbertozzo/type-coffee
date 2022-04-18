package usecase

import (
	"errors"

	"github.com/richardbertozzo/type-coffee/coffee"
	"github.com/richardbertozzo/type-coffee/coffee/entity"
)

type useCase struct {
	db coffee.Repository
}

// NewService returns a new coffee use case
func NewService(db coffee.Repository) coffee.UseCase {
	return useCase{
		db: db,
	}
}

func (u useCase) Create(c coffee.Coffee) error {
	link, err := entity.NewImageLink(c.Image)
	if err != nil {
		return err
	}
	cs, err := entity.NewCaracteristics(c.Caracteristics)
	if err != nil {
		return err
	}

	coff, err := entity.New(c.UUID, c.Name, c.Description, link, cs)
	if err != nil {
		return err
	}
	return u.db.Save(coff)
}

func (u useCase) GetByID(id string) (coffee.Coffee, error) {
	if id == "" {
		return coffee.Coffee{}, errors.New("id must be not blank")
	}

	c, err := u.db.GetByID(id)
	if err != nil {
		return coffee.Coffee{}, err
	}

	return convertCoffee(c), nil
}

func (u useCase) ListByCaracteristic(caStr string) (cos []coffee.Coffee, err error) {
	ca, err := entity.NewCaracteristic(caStr)
	if err != nil {
		return []coffee.Coffee{}, err
	}

	coffees, err := u.db.ListByCaracteristic(ca)
	if err != nil {
		return []coffee.Coffee{}, err
	}

	for _, co := range coffees {
		c1 := convertCoffee(co)
		cos = append(cos, c1)
	}

	return cos, nil
}

func (u useCase) List() (cos []coffee.Coffee, err error) {
	coffees, err := u.db.List()
	if err != nil {
		return []coffee.Coffee{}, err
	}

	for _, co := range coffees {
		c1 := convertCoffee(co)
		cos = append(cos, c1)
	}

	return cos, nil
}

func convertCoffee(c entity.Coffee) coffee.Coffee {
	casStr := []string{}
	for _, ca := range c.Caracteristics() {
		casStr = append(casStr, ca.String())
	}

	return coffee.Coffee{
		UUID:           c.ID,
		Name:           c.Name,
		Description:    c.Description,
		Image:          c.Image.String(),
		Caracteristics: casStr,
	}
}

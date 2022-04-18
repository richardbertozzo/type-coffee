package repository

import (
	"errors"

	"github.com/richardbertozzo/type-coffee/coffee"
	"github.com/richardbertozzo/type-coffee/coffee/entity"
)

type memoryDB struct {
	coffees map[string]entity.Coffee
}

// NewMemoryDB returns an in memory database implements the Coffee repository
func NewMemoryDB() coffee.Repository {
	m := map[string]entity.Coffee{}
	return memoryDB{
		coffees: m,
	}
}

func (m memoryDB) GetByID(id string) (entity.Coffee, error) {
	c := m.coffees[id]
	if c.IsZero() {
		return c, errors.New("Coffee not exists")
	}

	return c, nil
}

func (m memoryDB) Save(c entity.Coffee) error {
	m.coffees[c.ID] = c
	return nil
}

func (m memoryDB) ListByCaracteristic(c1 entity.Caracteristic) (cos []entity.Coffee, err error) {
	for _, c := range m.coffees {
		for _, carac := range c.Caracteristics() {
			if c1 == carac {
				cos = append(cos, c)
			}
		}
	}
	return
}

func (m memoryDB) List() (cos []entity.Coffee, err error) {
	for _, c := range m.coffees {
		cos = append(cos, c)
	}
	return
}

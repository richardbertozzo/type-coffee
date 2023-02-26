package repository

import (
	"errors"

	"github.com/richardbertozzo/type-coffee/coffee"
	"github.com/richardbertozzo/type-coffee/coffee/entity"
)

type memoryDB struct {
	coffees map[string]entity.Coffee
}

func (m memoryDB) ListByCaracteristic(characteristic coffee.Characteristic) ([]entity.Coffee, error) {
	//TODO implement me
	panic("implement me")
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

func (m memoryDB) List() (cos []entity.Coffee, err error) {
	for _, c := range m.coffees {
		cos = append(cos, c)
	}
	return
}

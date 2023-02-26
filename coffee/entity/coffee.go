package entity

import (
	"errors"
	"fmt"
)

// Coffee represent the entity of a Coffee
type Coffee struct {
	ID          string
	Name        string
	Image       Link
	Description string
}

func (c Coffee) String() string {
	return fmt.Sprintf("ID: %s - Name: %s", c.ID, c.Name)
}

// IsZero returns if coffee is nil/empty
func (c Coffee) IsZero() bool {
	return c.ID == ""
}

// New create a new coffee value object
func New(uuid, name, d string, l Link) (Coffee, error) {
	if name == "" {
		return Coffee{}, errors.New("name is blank")
	} else if uuid == "" {
		return Coffee{}, errors.New("UUID is null")
	}

	return Coffee{
		ID:          uuid,
		Name:        name,
		Image:       l,
		Description: d,
	}, nil
}

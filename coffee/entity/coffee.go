package entity

import (
	"errors"
	"fmt"
)

// Coffee represent the entity of a Coffee
type Coffee struct {
	ID             string
	Name           string
	Image          Link
	Description    string
	caracteristics []Caracteristic
}

func (c Coffee) String() string {
	return fmt.Sprintf("ID: %s - Name: %s", c.ID, c.Name)
}

// IsZero returns if coffee is nil/empty
func (c Coffee) IsZero() bool {
	return c.ID == ""
}

// Caracteristics returns the caracteristics of coffee
func (c Coffee) Caracteristics() []Caracteristic {
	return c.caracteristics
}

// New create a new coffee value object
func New(uuid, name, d string, l Link, c []Caracteristic) (Coffee, error) {
	if name == "" {
		return Coffee{}, errors.New("Name is blank")
	} else if uuid == "" {
		return Coffee{}, errors.New("UUID is null")
	} else if len(c) < 1 || len(c) > 4 {
		return Coffee{}, errors.New("Coffee must has more than 0 and less than 4 caracteristic")
	}

	return Coffee{
		ID:             uuid,
		Name:           name,
		Image:          l,
		Description:    d,
		caracteristics: c,
	}, nil
}

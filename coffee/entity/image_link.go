package entity

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// Link represent the entity of a Link
type Link struct {
	l string
}

func (l Link) String() string {
	return l.l
}

// IsZero returns if link is nil/empty
func (l Link) IsZero() bool {
	return l == Link{}
}

// NewImageLink create a new image link entity value object
func NewImageLink(s string) (Link, error) {
	u, err := url.ParseRequestURI(s)
	if err != nil {
		return Link{}, fmt.Errorf("Link invalid: %s", err.Error())
	}

	if !strings.Contains(u.Path, ".jpg") && !strings.Contains(u.Path, ".jpeg") && !strings.Contains(u.Path, ".png") {
		return Link{}, errors.New("Image link must be jpg, jpeg or png extensions")
	}

	// .jpg, .jpeg e .png
	return Link{
		l: u.String(),
	}, nil
}

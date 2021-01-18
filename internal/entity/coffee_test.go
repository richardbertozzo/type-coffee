package entity

import "testing"

func TestNewCoffee(t *testing.T) {
	t.Run("Must create a valid coffee", func(t *testing.T) {
		c := []Caracteristic{
			{
				Name: "carac 1",
			},
			{
				Name: "carac 2",
			},
		}
		c1, err := New("123", "Café Teste", c)
		if err != nil {
			t.Errorf("Must not return a error, but got %v", err)
		}
		t.Log(c1)
	})

	t.Run("Must create an invalid coffee", func(t *testing.T) {
		c := []Caracteristic{
			{
				Name: "carac 1",
			},
			{
				Name: "carac 2",
			},
		}
		c1, err := New("123", "", c)
		if err != errNameBlank {
			t.Errorf("Must return a error and must be blank name error, but got %v", err)
		}
		t.Log(c1)
	})

	t.Run("Must create an invalid coffee", func(t *testing.T) {
		c := []Caracteristic{
			{
				Name: "carac 1",
			},
			{
				Name: "carac 2",
			},
		}
		c1, err := New("123", "", c)
		if err != errNameBlank {
			t.Errorf("Must return a error and must be blank name error, but got %v", err)
		}
		t.Log(c1)
	})
}

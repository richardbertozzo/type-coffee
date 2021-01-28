package main

import (
	"log"

	"github.com/richardbertozzo/type-coffee/coffee"
	"github.com/richardbertozzo/type-coffee/coffee/repository"
	"github.com/richardbertozzo/type-coffee/coffee/usecase"
)

func main() {
	db := repository.NewMemoryDB()
	uc := usecase.NewService(db)

	c1 := coffee.Coffee{
		UUID:           "1234",
		Name:           "Cafe teste",
		Image:          "https://rollingstone.uol.com.br/media/_versions/godzilla-kingking-reprod-twitter-cortada_widelg.jpg",
		Description:    "dkajkdjajdsjadjsa",
		Caracteristics: []string{"fraco"},
	}
	err := uc.Create(c1)
	if err != nil {
		log.Fatal(err)
	}

	c2, err := uc.GetByID(c1.UUID)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Nome: %s - Link: %s", c2.Name, c2.Image)

	c3 := coffee.Coffee{
		UUID:           "1234",
		Name:           "Teste 2",
		Image:          "https://rollingstone.uol.com.br/media/_versions/godzilla-kingking-reprod-twitter-cortada_widelg.png",
		Description:    "",
		Caracteristics: []string{"jsdjasjdasjjdsasjk"},
	}
	err = uc.Create(c3)
	if err != nil {
		log.Fatal(err)
	}
}

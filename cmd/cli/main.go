package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/AlecAivazis/survey/v2"

	"github.com/richardbertozzo/type-coffee/coffee"
	"github.com/richardbertozzo/type-coffee/coffee/provider"
	"github.com/richardbertozzo/type-coffee/coffee/usecase"
	"github.com/richardbertozzo/type-coffee/infra/database"
)

func main() {
	var chatGPTKey, dbURL string
	flag.StringVar(&chatGPTKey, "CHAT_GPT_KEY", "", "Chat GPT API Key value")
	flag.StringVar(&dbURL, "DATABASE_URL", "", "Database URL value")
	flag.Parse()

	if chatGPTKey == "" {
		log.Fatal("CHAT_GPT_KEY ENV is required")
	}
	provider, err := provider.NewChatGPTProvider(chatGPTKey)
	if err != nil {
		log.Fatal(err)
	}

	var dbService coffee.Service
	if dbURL != "" {
		log.Println("database mode service enabled")
		dbPool, err := database.NewConnection(context.Background(), dbURL)
		if err != nil {
			log.Fatal(err)
		}
		dbService = provider.NewDatabaseService(dbPool)
	}

	uc := usecase.New(provider, dbService)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancelFunc()

	var inputCarac []string
	prompt := &survey.MultiSelect{
		Message:  "Select Up to 3 Characteristics",
		Options:  coffee.ListAllCharacteristic(),
		PageSize: 7,
	}

	err = survey.AskOne(prompt, &inputCarac, nil)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	bestCoffees, err := uc.GetBestCoffees(ctx, coffee.Filter{
		Characteristics: coffee.ConvertToCharacteristic(inputCarac),
	})
	if err != nil {
		log.Fatalf("error in get best coffee: %v", err)
	}

	// print options got from chat GPT openapi
	for i, opt := range bestCoffees.ChatGpt {
		fmt.Printf("Chat GPT %d option\n", i)
		fmt.Println(opt.Message)
	}

	// print options got from database data
	optsDB := bestCoffees.Database
	if optsDB != nil {
		for i, opt := range *optsDB {
			fmt.Printf("Database %d option\n", i)
			fmt.Println(opt.Message)
		}
	}
}

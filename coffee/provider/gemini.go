package provider

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	"github.com/richardbertozzo/type-coffee/coffee"
)

type GeminiClient struct {
	model *genai.GenerativeModel
}

func NewGeminiClient(apiKey string) (*GeminiClient, error) {
	ctx := context.Background()
	opt := option.WithAPIKey(apiKey)

	client, err := genai.NewClient(ctx, opt)
	if err != nil {
		return nil, err
	}

	model := client.GenerativeModel("gemini-1.5-flash")

	model.SetTemperature(1)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"

	return &GeminiClient{
		model: model,
	}, nil
}

func (c *GeminiClient) GetCoffeeOptionsByCharacteristics(ctx context.Context, filter coffee.Filter) ([]coffee.OptionProvider, error) {
	session := c.model.StartChat()
	session.History = []*genai.Content{}

	prompt := getPrompt(filter.Characteristics)
	log.Printf("prompt: %s", prompt)
	resp, err := session.SendMessage(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	if len(resp.Candidates) > 0 {
		opts := make([]coffee.OptionProvider, len(resp.Candidates[0].Content.Parts))
		for i, part := range resp.Candidates[0].Content.Parts {
			opts[i] = coffee.OptionProvider{
				Message: fmt.Sprintf("%v", part),
			}
		}

		return opts, nil
	}

	return nil, nil
}

func getPrompt(carats []coffee.Characteristic) string {
	promptTemplate := "recommend some good coffee options considering %s characteristics"

	var cs []string
	for _, c := range carats {
		cs = append(cs, string(c))
	}

	finalCharacteristics := strings.Join(cs, " and ")
	return fmt.Sprintf(promptTemplate, finalCharacteristics)
}

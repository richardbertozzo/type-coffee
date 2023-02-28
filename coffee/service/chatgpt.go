package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/richardbertozzo/type-coffee/coffee"
)

const (
	openaiApiUrl = "https://api.openai.com/v1/completions"
	model        = "text-davinci-003"
)

const promptTemplate = "Best %s Coffee"

type request struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float32 `json:"temperature"`
}

type completionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Model   string `json:"model"`
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

type openAIClient struct {
	http   *http.Client
	apiKey string
}

func NewChatGPTProvider(apiKey string) (coffee.Service, error) {
	return &openAIClient{
		http:   &http.Client{},
		apiKey: apiKey,
	}, nil
}

func getPrompt(carats []coffee.Characteristic) string {
	var cs []string
	for _, c := range carats {
		cs = append(cs, string(c))
	}

	finalCharacteristics := strings.Join(cs, " ")
	return fmt.Sprintf(promptTemplate, finalCharacteristics)
}

func (c *openAIClient) GetCoffeeOptionsByCharacteristics(ctx context.Context, filter coffee.Filter) ([]coffee.OptionProvider, error) {
	prompt := getPrompt(filter.Characteristics)
	fmt.Println(prompt)

	reqBody := request{
		Model:       model,
		Prompt:      prompt,
		Temperature: 0.2,
		MaxTokens:   100,
	}

	payload, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request payload: %s", err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, openaiApiUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request:: %s", err.Error())
	}
	defer resp.Body.Close()

	// Parse the response payload
	var respBody completionResponse
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, fmt.Errorf("error parsing response payload: %s", err.Error())
	}

	opts := make([]coffee.OptionProvider, len(respBody.Choices))
	for i, choice := range respBody.Choices {
		opts[i] = coffee.OptionProvider{
			Message: choice.Text,
		}
	}

	return opts, nil
}

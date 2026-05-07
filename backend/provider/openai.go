package provider

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIClient struct {
	client openai.Client
	model  string
}

func NewOpenAIClient(apiKey, baseURL, model string) *OpenAIClient {
	return &OpenAIClient{
		client: openai.NewClient(
			option.WithAPIKey(apiKey),
			option.WithBaseURL(baseURL),
		),
		model: model,
	}
}

func (c *OpenAIClient) Chat(ctx context.Context, prompt string) (string, error) {
	resp, err := c.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: c.model,
	})
	if err != nil {
		return "", fmt.Errorf("LLM API call failed: %w", err)
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices in LLM response")
	}
	return resp.Choices[0].Message.Content, nil
}

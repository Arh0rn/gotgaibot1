package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Arh0rn/gotgaibot1/internal/llm"
	"io"
	"log/slog"
	"net/http"
)

func (c *OpenAIClient) GenerateResponse(ctx context.Context, history []llm.Message) (string, error) {
	reqBody := ChatRequest{
		Model:       c.Model,
		Messages:    history,
		MaxTokens:   c.MaxTokens,
		Temperature: c.Temperature,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", c.BaseUrl+"/chat/completions", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	slog.Info(string(b))

	var chatResponse ChatResponse
	if err := json.Unmarshal(b, &chatResponse); err != nil {
		return "", err
	}
	if len(chatResponse.Choices) > 0 {
		return chatResponse.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no answers found in the response")
}

func (c *OpenAIClient) GetLegend() string {
	return c.Legend
}

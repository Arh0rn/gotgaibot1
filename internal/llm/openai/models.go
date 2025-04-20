package openai

import (
	"github.com/Arh0rn/gotgaibot1/internal/llm"
	"github.com/Arh0rn/gotgaibot1/pkg/config"
	"net/http"
)

type OpenAIClient struct {
	ApiKey      string
	BaseUrl     string
	Model       string
	Temperature float64
	MaxTokens   int
	httpClient  *http.Client
	Legend      string
}

func New(cfg config.LLMConfig) *OpenAIClient {
	return &OpenAIClient{
		ApiKey:      cfg.APIKey,
		BaseUrl:     cfg.BaseUrl,
		Model:       cfg.Model,
		Temperature: cfg.Temperature,
		MaxTokens:   cfg.MaxTokens,
		Legend:      cfg.Legend,
		httpClient:  &http.Client{},
	}
}

type ChatRequest struct {
	Model       string        `json:"model"`
	Messages    []llm.Message `json:"messages"`
	MaxTokens   int           `json:"max_tokens"`
	Temperature float64       `json:"temperature"`
}

type ChatResponse struct {
	Choices []struct {
		Message llm.Message `json:"message"`
	} `json:"choices"`
}

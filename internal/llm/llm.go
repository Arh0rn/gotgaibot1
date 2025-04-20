package llm

import "context"

type LLM interface {
	GenerateResponse(ctx context.Context, history []Message) (string, error)
	GetLegend() string
}

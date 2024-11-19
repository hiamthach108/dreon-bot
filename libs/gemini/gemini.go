package gemini

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var (
	ctx        = context.Background()
	model      = "gemini-1.5-flash"
	embedModel = "text-embedding-004"
)

type GenAIGemini struct {
	IGenAIGemini
	model      *genai.GenerativeModel
	embedModel *genai.EmbeddingModel
}

type IGenAIGemini interface {
	NewGenAIGemini(apiKey string) (*GenAIGemini, error)
	GenerateContent(message string) (resp string, err error)
	EmbedContent(message string) (vector []float32, err error)
}

func NewGenAIGemini(apiKey string) (*GenAIGemini, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	// defer client.Close()

	model := client.GenerativeModel(model)
	em := client.EmbeddingModel(embedModel)

	return &GenAIGemini{
		model:      model,
		embedModel: em,
	}, nil
}

func (g *GenAIGemini) GenerateContent(message string) (resp string, err error) {
	msg := genai.Text(message)
	reply, err := g.model.GenerateContent(ctx, genai.Text(msg))
	if err != nil {
		return "", err
	}

	resp = ""
	for _, c := range reply.Candidates {
		for _, p := range c.Content.Parts {
			resp += fmt.Sprintf("%s", p)
		}
	}

	return resp, nil
}

func (g *GenAIGemini) EmbedContent(message string) (vector []float32, err error) {
	reply, err := g.embedModel.EmbedContent(ctx, genai.Text(message))
	if err != nil {
		return nil, err
	}

	return reply.Embedding.Values, nil
}

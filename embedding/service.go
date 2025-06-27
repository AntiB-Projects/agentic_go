package embedding

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type Embedder struct {
	client *genai.Client
	model  string
}

// NewEmbedder is the exported constructor
func NewEmbedder(ctx context.Context, model string) (*Embedder, error) {
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &Embedder{client: client, model: model}, nil
}

// Embed returns the first embedding vector for the given text
func (e *Embedder) Embed(ctx context.Context, text string) ([]float64, error) {
	contents := []*genai.Content{
		genai.NewContentFromText(text, genai.RoleUser),
	}
	resp, err := e.client.Models.EmbedContent(ctx, e.model, contents, nil)
	if err != nil {
		return nil, err
	}
	if len(resp.Embeddings) == 0 {
		return nil, fmt.Errorf("no embeddings returned")
	}
	return resp.Embeddings[0].Embedding, nil
}

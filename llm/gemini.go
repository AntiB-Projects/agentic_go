package llm 

import (
	"fmt"
	"context"

	genai "google.golang.org/genai"
)

type GeminiClient struct {
	client 		*genai.Client	
}

func NewGeminiClient(ctx context.Context, apiKey string) (*GeminiClient, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create genai client: %w", err)
	}

	return &GeminiClient{client: client}, nil
}

func (c *GeminiClient) Chat(ctx context.Context, req ChatRequest) (ChatMessage, error) {
	
	// 1. Setting system configuration
	config := &genai.GenerateContentConfig{
		SystemInstruction:	&genai.Content{Parts : []*genai.Part{{Text: req.SystemPrompt}},},
		Temperature:      	&req.Temperature,
	}

	// 2. Setting history
	history := []*genai.Content{}
	for _, msg := range req.History {
		var currentRole string

		if msg.Role == "assistant" {
			currentRole = genai.RoleModel
		} else {
			currentRole = genai.RoleUser
		}

		history = append(history, &genai.Content{
			Role:	currentRole,
			Parts:	[]*genai.Part{{Text: msg.Content}},
		})
	}

	// 3. Creating chat with gemini
	chat, err := c.client.Chats.Create(ctx, req.Model, config, history)
	if err != nil {
		return ChatMessage{}, fmt.Errorf("gemini sdk: failed to send message: %w", err)
}

	// 4. Making API-Call with new chat
	result, err := chat.SendMessage(ctx, genai.Part{Text: req.Message})
	if err != nil {
		return ChatMessage{}, fmt.Errorf("gemini sdk: failed to send message: %w", err)
	}

	// 5. Translate the SDK's response back to our internal ChatResponse format.
	if len(result.Candidates) == 0 || result.Candidates[0].Content == nil {
		return ChatMessage{}, fmt.Errorf("gemini sdk: no candidates returned")
	}

	// The SDK gives us parts, we need to assemble them into a single string.
	var responseContent string
	candidate := result.Candidates[0]
	for _, part := range candidate.Content.Parts {
		if part != nil {
			responseContent += part.Text
		}
	}
	
	response := ChatMessage{
		Role:    "assistant",
		Content: responseContent,
	}
	return response, nil
}

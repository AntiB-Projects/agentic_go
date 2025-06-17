package llm

import "context"

type ChatMessage struct {
	Role 			string // 'user' or 'assistant'
	Content			string
}

type ChatRequest struct {
	Model 			string
	SystemPrompt	string
	Message			string
	History 		[]ChatMessage
	Temperature		float32
}

type LLM interface {
	Chat(ctx context.Context, req ChatRequest) (ChatMessage, error)
} 

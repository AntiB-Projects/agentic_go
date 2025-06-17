package agent

import (
	"context"
	"fmt"

	"github.com/AntiB-Projects/agentic_go/llm"

)

type Agent struct {
	llm 			llm.LLM
	modelName		string
	systemPrompt	string
	history			[]llm.ChatMessage
	temperature		float32
}

// Creates new Agent with empty history
func NewAgent(llmProvider llm.LLM, modelName, systemPrompt string) *Agent {
	return &Agent{
		llm:			llmProvider,
		modelName:		modelName,
		systemPrompt:	systemPrompt,
		history:		[]llm.ChatMessage{},
	}
}

func (a *Agent) Respond(ctx context.Context, userMessage string) (string, error) {
	// create request 
	req := llm.ChatRequest{
		Model: 			a.modelName,
		SystemPrompt:	a.systemPrompt,
		Message:		userMessage,
		History:		a.history,
		Temperature:	a.temperature,
	}
	
	// call LLM for response
	response, err := a.llm.Chat(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to get response from LLM: %w", err)
	}

	// append new messages to agents` history
	a.history = append(a.history, llm.ChatMessage{Role: "user", Content: userMessage})
	a.history = append(a.history, response)

	return response.Content, nil
}

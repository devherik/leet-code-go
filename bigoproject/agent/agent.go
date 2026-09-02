package agent

import (
	"context"
	"encoding/json"
	"fmt"
)

type AgentDecision struct {
	Action    string          `json:"action"`
	Arguments json.RawMessage `json:"arguments"`
	Reasoning string          `json:"reasoning"`
}

type QueryDBArgs struct {
	SQL string `json:"sql"`
}

type LLMClient interface {
	GenerateJSON(ctx context.Context, prompt string, schema string) ([]byte, error)
}

type AgentOrchestrator struct {
	llm LLMClient
}

func (a *AgentOrchestrator) Step(ctx context.Context, systemPrompt string) error {
	rawOutput, err := a.llm.GenerateJSON(ctx, systemPrompt, `{"type": "object", "properties": {"action": {"type": "string"}, "arguments": {"type": "object"}, "reasoning": {"type": "string"}}, "required": ["action", "arguments", "reasoning"]}}`)
	if err != nil {
		return fmt.Errorf("llm generation failed: %w", err)
	}

	var decision AgentDecision
	if err := json.Unmarshal(rawOutput, &decision); err != nil {
		return fmt.Errorf("llm broke output: %w", err)
	}
	switch decision.Action {
	case "query_db":
		var args QueryDBArgs
		if err := json.Unmarshal(decision.Arguments, &args); err != nil {
			return fmt.Errorf("invalid tool args: %w", err)
		}
		return nil
	case "respond":
		return nil
	default:
		return fmt.Errorf("unrecognized: %s", decision.Action)
	}
}

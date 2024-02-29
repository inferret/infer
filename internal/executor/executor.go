package executor

import (
	"encoding/json"
	"fmt"

	"github.com/inferret/infer/internal/api"
	"github.com/inferret/infer/internal/parser"
	"github.com/sashabaranov/go-openai" // Assuming this is where ChatCompletionMessage is defined.
)

// jsonResponse is the expected JSON structure of the response.
type jsonResponse struct {
	Assertion bool `json:"assertion"`
}

type Executor struct {
	Client  *api.OpenAIWrapper
	Verbose bool
}

func NewExecutor(client *api.OpenAIWrapper, verbose bool) *Executor {
	return &Executor{
		Client:  client,
		Verbose: verbose,
	}
}

// Infer: Execute
func (e *Executor) Execute(inference parser.Inference, code string) (bool, error) {
	successCount := 0
	totalCount := inference.Count

	for i := 0; i < totalCount; i++ {
		var tag = inference.Tag_Name
		messages := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Here is a code block tagged as [" + tag + "]. Please analyze the following code: \n\n" + code,
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Based on the code analysis, answer the following question with a JSON-formatted boolean response in the format: ```{\"assertion\": true}```.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Is the following assertion about the code true? " + inference.Assertion,
			},
		}

		// Adjusted call to CreateChatCompletion without context.Context and with correct parameter order.
		completion, err := e.Client.CreateChatCompletion(messages, inference.Model, inference.MaxTokens, inference.Temperature)
		if err != nil {
			return false, fmt.Errorf("chat completion error: %w", err)
		}

		// Check if there are any choices in the completion response.
		if completion.Choices == nil || len(completion.Choices) == 0 {
			return false, fmt.Errorf("no choices in completion response")
		}

		// Unmarshal the JSON response from the first choice's message content.
		var resp jsonResponse
		if err := json.Unmarshal([]byte(completion.Choices[0].Message.Content), &resp); err != nil {
			return false, fmt.Errorf("failed to unmarshal response JSON: %w", err)
		}

		if resp.Assertion {
			successCount++
		}
	}

	// Calculate the success rate
	successRate := float64(successCount) / float64(totalCount)
	// Determine if the success rate meets the threshold
	if successRate < inference.Threshold {
		if e.Verbose || successRate <= inference.Threshold {
			fmt.Printf("Inference failed: %s. Success rate: %.2f%% (Threshold: %.2f%%)\n", inference.Assertion, successRate*100, inference.Threshold*100)
		}
		return false, nil
	}

	if e.Verbose {
		fmt.Printf("Inference successful: %s. Success rate: %.2f%% (Threshold: %.2f%%)\n", inference.Assertion, successRate*100, inference.Threshold*100)
	}

	return true, nil
}

//InferEnd: Execute

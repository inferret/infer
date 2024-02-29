// Infer: OpenAI client
package api

import (
	"context"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

// OpenAIWrapper encapsulates the OpenAI client and provides methods for executing inferences.
type OpenAIWrapper struct {
	Client  *openai.Client
	Verbose bool // Add a verbosity flag
}

// NewOpenAIWrapper initializes a new OpenAI client with the provided API key and verbosity setting.
func NewOpenAIWrapper(apiKey string, baseUrl string, verbose bool) *OpenAIWrapper {
	if apiKey == "" {
		apiKey = os.Getenv("OPENAI_API_KEY") // Fallback to environment variable if not provided
	}
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = baseUrl
	client := openai.NewClientWithConfig(config)

	return &OpenAIWrapper{
		Client:  client,
		Verbose: verbose,
	}
}

func (o *OpenAIWrapper) CreateChatCompletion(messages []openai.ChatCompletionMessage, model string, maxTokens int, temperature float64) (openai.ChatCompletionResponse, error) {
	// Create an instance of ChatCompletionResponseFormat with Type set to JSON object
	responseFormat := openai.ChatCompletionResponseFormat{
		Type: openai.ChatCompletionResponseFormatTypeJSONObject,
	}

	chatCompletionParams := openai.ChatCompletionRequest{
		Model:          model,
		Messages:       messages,
		MaxTokens:      maxTokens,
		Temperature:    float32(temperature),
		ResponseFormat: &responseFormat, // Set the response format to JSON object
	}

	if o.Verbose {
		log.Printf("Creating chat completion with parameters: Model=%s, MaxTokens=%d, Temperature=%f, ResponseFormat=%s", model, maxTokens, temperature, chatCompletionParams.ResponseFormat)
		for i, msg := range messages {
			// Only print the User messages, not the system messages
			if msg.Role == openai.ChatMessageRoleUser {
				log.Printf("Message %d: Role=%s, Content=%s", i, msg.Role, msg.Content)
			}
		}
	}

	response, err := o.Client.CreateChatCompletion(context.Background(), chatCompletionParams)
	if err != nil {
		return openai.ChatCompletionResponse{}, err
	}

	return response, nil
}

// InferEnd: OpenAI client

package lib

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func Chat() (string, error) {
	// Load .env file
	godotenv.Load()

	// Prompt user and flush stdout immediately
	fmt.Print("? Describe your changes in a few words (e.g., 'Fixed bug in Auth'): ")
	os.Stdout.Sync() // Fixed invisible prompt

	// Scanner reads user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading input: %v", err)
	}
	userInput := scanner.Text()

	prompt := fmt.Sprintf("Input: %s\n", userInput)

	// Loads API Key from OpenRouter.ai (.env)
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Invalid API Key!")
	}

	// Configure client for OpenRouter
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://openrouter.ai/api/v1"

	client := openai.NewClientWithConfig(config)

	// OpenAI
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			// Use DeepSeek V3 model
			Model: "deepseek/deepseek-chat:free",
			// Message with prompt prefix
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: os.Getenv("Prompt_prefix") + prompt,
				},
			},
		})

	// Error check
	if err != nil {
		return "", fmt.Errorf("completion error: %v", err)
	}

	// Returns response
	return resp.Choices[0].Message.Content, nil
}

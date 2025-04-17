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

	PrintWelcome()

	// Prompt user and flush stdout immediately
	fmt.Print("? Describe your changes in a few words (e.g., 'Fixed bug in Auth'): ")
	os.Stdout.Sync() // Fixed invisible prompt

	// Scanner reads user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// Scanner error check
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading input: %v", err)
	}

	// Saves user input
	userInput := scanner.Text()
	prompt := fmt.Sprintf("Input: %s\n", userInput)

	println("\n🔄 Generating commit message...\n")

	// Loads API Key from OpenRouter.ai (.env)
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("\n❌ Invalid API Key")
	}

	// Configure client for OpenRouter
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://openrouter.ai/api/v1"
	client := openai.NewClientWithConfig(config)

	// OpenAI connection
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
		println("\n❌ Failed to generate commit message.\n")
		println("Possible reasons:\n\t- No internet connection\n\t- API rate limit exceeded\n")
		println("Press [Ctrl + C] to exit, and try again.")
	}

	// Returns response
	return resp.Choices[0].Message.Content, nil
}

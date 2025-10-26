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

func Generate() (string, error) {
	// Load .env file
	godotenv.Load()

	PrintWelcome()

	// Prompt user and flush stdout immediately
	fmt.Print("? Describe your changes in a few words (e.g., 'Fixed bug in Auth'): ")
	os.Stdout.Sync()

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

	println("📄 Generating commit message...")

	// Loads API Key from Groq (.env)
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("\n❌ Invalid API Key")
	}

	// Configure client for Groq
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://api.groq.com/openai/v1" // Changed to Groq endpoint
	client := openai.NewClientWithConfig(config)

	// Groq API connection
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			// Use Llama 3.3 70B - Fast and high quality
			Model: "llama-3.3-70b-versatile",
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
		println("👋 Exiting...")
		return "", err
	}

	// Returns response
	return resp.Choices[0].Message.Content, nil
}

package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

// Default shared API key - This is INTENTIONAL and no harm can be done with this key
const defaultAPIKey = "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"

const defaultPromptPrefix = `You are a Conventional Commit expert. Generate a commit message that follows strict conventional commit standards. Format as:

[type]([scope]): [imperative-mood-subject-under-50-chars]

Structure:
1. Type must be one of: feat, fix, chore, refactor, docs, style, test, perf, ci, build, revert
2. Scope should indicate the module/component (e.g., auth, database, ui)
3. Subject must:
   - Start with lowercase verb (fix, add, remove)
   - Use imperative mood ('Add' not 'Added')
   - Contain no ending punctuation
   - Stay under 50 characters

Example Transformation:
Input: 'fixed bug in login where user can make acc. without a passwrd'
Output: 'fix(auth): resolve login validation bug for empty passwords'

Current User Input to Transform:
`

func GenerateFromDescription(description string) (string, error) {
	// Load .env file (optional - for users who want custom keys)
	godotenv.Load()

	// Build prompt
	prompt := fmt.Sprintf("Input: %s\n", description)

	// Use env var if present, otherwise use default shared key
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		apiKey = defaultAPIKey
	}

	// Use env prompt prefix if present, otherwise use default
	promptPrefix := os.Getenv("Prompt_prefix")
	if promptPrefix == "" {
		promptPrefix = defaultPromptPrefix
	}

	// Configure client for Groq
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://api.groq.com/openai/v1"
	client := openai.NewClientWithConfig(config)

	// Groq API connection
	resp, err := client.CreateChatCompletion(
		context.Background(),
		// Use openAI
		openai.ChatCompletionRequest{
			Model: "llama-3.3-70b-versatile",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: promptPrefix + prompt,
				},
			},
		})

	if err != nil {
		return "", err
	}

	// Returns response
	return resp.Choices[0].Message.Content, nil
}

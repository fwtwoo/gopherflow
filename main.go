package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func printWelcome() {
	// Print welcome message
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                Welcome to GopherFlow!                    ║")
	fmt.Println("║  AI-generated, industry-standard commit messages.        ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Print("\n")
}

// func printError() {
// 	// Print API error message
// 	fmt.Println("❌ Failed to generate commit message.")
// 	fmt.Print("\n")
// 	fmt.Println("Possible reasons:\n- No internet connection\n- API rate limit exceeded")
// 	fmt.Print("\n")
// }

func main() {
	// Load .env file
	env := godotenv.Load()

	// Error handling for .env file
	if env != nil {
		log.Fatal("Error loading .env file")
	}

	// Get API Key
	apiKey := os.Getenv("API_KEY")

	// Print welcome message
	printWelcome()

	// Scanner saves input using Stdin (standard input)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// Save input
	userInput := scanner.Text()
	fmt.Println("\nYour input:", userInput)
	fmt.Printf("API Key: %s\n", apiKey)
}

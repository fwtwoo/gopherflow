package main

import (
	"fmt"
)

func printWelcome() {
	// Print welcome message
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                Welcome to GopherFlow!                    ║")
	fmt.Println("║       AI-generated, high-quality commit messages.        ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Print("\n")
}

func printError() {
	// Print API error message
	fmt.Println("❌ Failed to generate commit message.")
	fmt.Print("\n")
	fmt.Println("Possible reasons:\n- No internet connection\n- API rate limit exceeded")
	fmt.Print("\n")
}

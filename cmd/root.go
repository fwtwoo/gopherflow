package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/ekrlstd/summit/lib"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "summit [description]",
	Short: "Generate conventional commit messages",
	Long:  `GopherFlow generates high-quality conventional commit messages using AI.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Join all args into description
		description := strings.Join(args, " ")

		// Generate commit message
		res, err := lib.GenerateFromDescription(description)

		if err != nil {
			fmt.Println("⚠️ Failed to generate commit message")
			fmt.Println("Check your internet connection")
			os.Exit(1)
		}

		// Output with lightning bolt (hell yeah)
		fmt.Printf("⚡%s\n", res)

		// Copy to clipboard
		err = clipboard.WriteAll(res)
		if err == nil {
			fmt.Println("[Copied to clipboard]")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// No flags needed for minimal UX
}

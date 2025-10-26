package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd is kept for backwards compatibility but hidden
var generateCmd = &cobra.Command{
	Use:    "generate",
	Hidden: true,
	Short:  "Legacy command - use 'summit <description>' instead",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("⚠️ This command is deprecated")
		fmt.Println("Use: summit <description>")
		fmt.Println("Example: summit fixed login bug")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

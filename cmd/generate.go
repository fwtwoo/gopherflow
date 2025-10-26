package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd is kept for backwards compatibility but hidden
var generateCmd = &cobra.Command{
	Use:    "generate",
	Hidden: true,
	Short:  "Legacy command - use 'gopherflow <description>' instead",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("⚠️  This command is deprecated")
		fmt.Println("Use: gopherflow <description>")
		fmt.Println("Example: gopherflow fixed login bug")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

// Auto-generated with Cobra-CLI
package cmd

import (
	"fmt"
	"strings"

	"github.com/ebbekarlstad/gopherflow/lib"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Command to run the application.",

	Run: func(cmd *cobra.Command, args []string) {
		res, err := lib.Generate()
		// Check if response contains "❌" (means Deepseek gen. error)
		contains := strings.Contains(res, "❌")

		if err != nil {
			return
		}
		if contains {
			fmt.Println("")
			fmt.Println(res)
		} else {
			fmt.Println("\n✅ Commit message generated successfully: ")
			fmt.Println("──────────────────────────────────────────────")
			fmt.Println("\n", res)
			fmt.Println("\n──────────────────────────────────────────────")
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

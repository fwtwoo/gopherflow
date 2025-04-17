// Auto-generated with Cobra-CLI
package cmd

import (
	"fmt"
	"log"

	"github.com/ebbekarlstad/gopherflow/lib"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "run",
	Short: "Command to run the application.",

	Run: func(cmd *cobra.Command, args []string) {
		res, err := lib.Chat()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("> Response: ")
		fmt.Println("")
		fmt.Println(res)
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

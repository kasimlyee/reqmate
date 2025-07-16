package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute API test suite",
	Long:  `Run a collection of API tests defined in configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running test suite...")
		// Implementation will be added in Phase 2
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
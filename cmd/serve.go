package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start interactive TUI mode",
	Long:  `Launch the interactive terminal user interface`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting interactive mode...")
		// Implementation will be added in Phase 4
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
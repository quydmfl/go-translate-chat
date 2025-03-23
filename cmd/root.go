package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-translate-chat",
	Short: "A WebSocket chat server with private, group, and broadcast chat.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Go WebSocket Chat!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

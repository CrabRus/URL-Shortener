package cmd

import (
	"URLShorter/history"

	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Show URL shortening history",
	Run: func(cmd *cobra.Command, args []string) {
		history.Print()
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}

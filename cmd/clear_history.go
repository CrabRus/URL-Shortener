package cmd

import (
	"URLShorter/history"
	"fmt"

	"github.com/spf13/cobra"
)

var clearHistoryCmd = &cobra.Command{
	Use:   "clear-history",
	Short: "Clear URL shortening history",
	Run: func(cmd *cobra.Command, args []string) {
		err := history.Clear()
		if err != nil {
			fmt.Println("❌ Error clearing history:", err)
			return
		}
		fmt.Println("✅ History cleared successfully")
	},
}

func init() {
	rootCmd.AddCommand(clearHistoryCmd)
}

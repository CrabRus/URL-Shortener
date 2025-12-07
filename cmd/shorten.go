package cmd

import (
	"URLShorter/api"
	"URLShorter/history"
	"URLShorter/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	service string
	check   bool
)

var shortenCmd = &cobra.Command{
	Use:   "shorten [url]",
	Short: "Shorten a URL",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		original := utils.NormalizeURL(args[0])

		if check && !utils.CheckAvailability(original) {
			fmt.Println("⚠️ Warning: URL may be unreachable")
		}

		fmt.Println("📡 Sending request to API...")

		short, serviceName, err := api.ShortenByName(original, service)
		if err != nil {
			fmt.Println("❌", err)
			return
		}

		utils.ShowResult(original, short)
		history.Add(original, short, serviceName)
	},
}

func init() {
	shortenCmd.Flags().StringVarP(&service, "service", "s", "tinyurl",
		"URL shortening service: tinyurl or isgd")

	shortenCmd.Flags().BoolVarP(&check, "check", "c", false,
		"Check if URL is reachable before shortening")

	rootCmd.AddCommand(shortenCmd)
}

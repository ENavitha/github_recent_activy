/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github-recent-activity/github"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-recent-activity  <github username>",
	Short: "Fetch recent GitHub activity for a user",
	Long:  "fetching recent activity in github enter  github username",
	Args:  cobra.ExactArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(" Starting GitHub activity fetch...") // Debug message

		if len(args) < 1 {
			fmt.Println(" No username provided. Exiting.")
			return
		}

		username := args[0]
		fmt.Printf("🔍 Fetching activity for: %s\n", username)

		github.FetchGitHubActivity(username) // Call your function

		fmt.Println(" Done!")

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.github-recent-activity.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

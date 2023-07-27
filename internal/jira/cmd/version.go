package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of jira-cli",
	Long:  "All software has version. This is jira-cli's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 0.0.1 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

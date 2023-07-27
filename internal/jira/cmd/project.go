package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/iyuuya/jira/internal/client"
)

var projectRootCmd = &cobra.Command{
	Use: "project",
}

var projectListCmd = &cobra.Command{
	Use: "ls",
	Run: func(cmd *cobra.Command, args []string) {
		cl, err := client.NewClient(cmd.Context())
		if err != nil {
			log.Fatal(err)
		}

		ps, _, err := cl.Project.GetList()
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range *ps {
			fmt.Printf("%s: %s\n", v.Key, v.Name)
		}
	},
}

func init() {
	projectRootCmd.AddCommand(projectListCmd)
	rootCmd.AddCommand(projectRootCmd)
}

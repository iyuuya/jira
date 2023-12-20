package cmd

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/iyuuya/jira/internal/client"
)

var issueRootCmd = &cobra.Command{
	Use: "issue",
}

var issueListCmd = &cobra.Command{
	Use: "ls [PROJECT_KEY]",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := client.NewClient(cmd.Context())
		if err != nil {
			log.Fatal(err)
		}

		if len(args) == 0 {
			return errors.New("Need Project Key")
		}

		query := fmt.Sprintf("project = %s", args[0])
		if len(args) > 1 {
			for _, v := range args[1:len(args)] {
				query += " AND " + v
			}
		}

		issues, _, err := cl.Issue.Search(query, nil)
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range issues {
			fmt.Printf("%s: %s\n", v.Key, v.Fields.Summary)
		}
		return nil
	},
}

func init() {
	issueRootCmd.AddCommand(issueListCmd)
	rootCmd.AddCommand(issueRootCmd)
}

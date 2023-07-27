package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/iyuuya/jira/internal/client"
	"github.com/iyuuya/jira/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "jira",
	Short: "JIRA Command Line Interface",
}

func Execute() error {
	ctx := context.Background()
	defer ctx.Done()

	conf, err := config.LoadFromEnv()
	if err != nil {
		return err
	}

	ctx = withJiraContext(ctx, conf)
	rootCmd.SetContext(ctx)

	return rootCmd.Execute()
}

func withJiraContext(ctx context.Context, conf *config.Configuration) context.Context {
	ctx = context.WithValue(ctx, client.JiraHostKey, conf.JiraHost)
	ctx = context.WithValue(ctx, client.JiraUsernameKey, conf.JiraUsername)
	return context.WithValue(ctx, client.JiraPasswordKey, conf.JiraAPIToken)
}

package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/iyuuya/go/exitcode"

	"github.com/iyuuya/jira/internal/client"
	"github.com/iyuuya/jira/internal/config"
)

func main() {
	os.Exit(int(realMain()))
}

func realMain() exitcode.ExitCode {
	conf, err := config.LoadFromEnv()
	if err != nil {
		fmt.Printf("Failed to load config: %s\n", err)
		return exitcode.ExitError
	}

	ctx := context.Background()
	defer ctx.Done()
	ctx = withJiraContext(ctx, conf)

	cl, err := client.NewClient(ctx)

	q := strings.Join(os.Args[1:len(os.Args)], " ")

	issues, _, err := cl.Issue.Search(q, nil)
	if err != nil {
		fmt.Printf("Failed to search: %s\n", err)
		return exitcode.ExitError
	}

	for _, v := range issues {
		fmt.Printf("%s\n", v.Key)
	}

	return exitcode.ExitOK
}

func withJiraContext(ctx context.Context, conf *config.Configuration) context.Context {
	ctx = context.WithValue(ctx, client.JiraHostKey, conf.JiraHost)
	ctx = context.WithValue(ctx, client.JiraUsernameKey, conf.JiraUsername)
	return context.WithValue(ctx, client.JiraPasswordKey, conf.JiraAPIToken)
}

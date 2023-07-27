package client

import (
	"context"

	jira "github.com/andygrunwald/go-jira"
)

type ContextKey string

const JiraHostKey ContextKey = "jira_host"
const JiraUsernameKey ContextKey = "jira_username"
const JiraPasswordKey ContextKey = "jira_password"

func NewClient(ctx context.Context) (*jira.Client, error) {
	host := ctx.Value(JiraHostKey).(string)
	username := ctx.Value(JiraUsernameKey).(string)
	password := ctx.Value(JiraPasswordKey).(string)
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}

	return jira.NewClient(tp.Client(), host)
}

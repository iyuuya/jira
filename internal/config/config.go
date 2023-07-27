package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	JiraHost     string `envconfig:"JIRA_HOST"`
	JiraUsername string `envconfig:"JIRA_USERNAME"`
	JiraAPIToken string `envconfig:"JIRA_API_TOKEN"`
}

func LoadFromEnv() (*Configuration, error) {
	c := &Configuration{}
	err := envconfig.Process("", c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

package main

import (
	"fmt"
	"log"
	"os"

	jira "github.com/andygrunwald/go-jira"
	"github.com/iyuuya/go/exitcode"

	"github.com/iyuuya/jira/internal/cmd"
)

func main() {
	os.Exit(int(realMain()))
}

func realMain() exitcode.ExitCode {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		return exitcode.ExitError
	}

	return exitcode.ExitOK
}

func getProjects() {
	tp := jira.BasicAuthTransport{
		Username: os.Getenv("JIRA_USERNAME"),
		Password: os.Getenv("JIRA_API_TOKEN"),
	}

	client, err := jira.NewClient(tp.Client(), os.Getenv("JIRA_HOST"))
	if err != nil {
		log.Fatal(err)
	}

	ps, _, err := client.Project.GetList()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range *ps {
		fmt.Printf("%s: %s\n", v.Key, v.Name)
	}

}

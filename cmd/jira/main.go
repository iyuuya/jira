package main

import (
	"fmt"
	"os"

	"github.com/iyuuya/go/exitcode"

	"github.com/iyuuya/jira/internal/jira/cmd"
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

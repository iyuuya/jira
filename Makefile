export GO111MODULE=on

.PHONY: run
run:
	@go run github.com/iyuuya/jira/cmd/jira

bin/jira:
	@go build -o bin/jira github.com/iyuuya/jira/cmd/jira

export GO111MODULE=on

.PHONY: build
build: bin/jira bin/jql

.PHONY: run
run:
	@go run github.com/iyuuya/jira/cmd/jira

bin/jira:
	@go build -o bin/jira github.com/iyuuya/jira/cmd/jira

bin/jql:
	@go build -o bin/jql github.com/iyuuya/jira/cmd/jql

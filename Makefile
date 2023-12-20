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

install: install/jira install/jql

install/jira:
	@go install github.com/iyuuya/jira/cmd/jira
install/jql:
	@go install github.com/iyuuya/jira/cmd/jql

## Configuration

```shell
export JIRA_HOST=https://hoge.atlassian.net
export JIRA_USERNAME=your-email@example.com
export JIRA_API_TOKEN=personal-api-token
```

## Example

```
bin/jira issue ls {YOUR_PROJECT} 'issuetype = Story and "Story Points[Number]" is EMPTY order by rank asc'                                                                                                                                                                                          [~/go/src/github.com/iyuuya/jira]
```

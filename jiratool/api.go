package jiratool

import (
	"log"

	jira "github.com/andygrunwald/go-jira"
)

var jiraClient *jira.Client

func Init(username, apiToken, baseUrl string) {
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: apiToken,
	}

	cli, err := jira.NewClient(tp.Client(), baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	jiraClient = cli
}

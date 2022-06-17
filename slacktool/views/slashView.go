package views

import (
	"embed"
	"encoding/json"
	"io/ioutil"

	"github.com/slack-go/slack"
)

const (
	// Define Action_id as constant so we can refet to them in the controller
	AssigneeActionId = "assignee-action-id"
	ReporterActionId = "repoter-action-id"

	CreateJiraTicketActionId = "create-jira-ticket-action-id"

	IssueTypeActionId = "issue-type-action-id"
)

//go:embed jiraAssets/*
var slashCommandAssets embed.FS

func CreateJiraInfoView(ticketTitle string) []slack.Block {
	// we need a stuct to hold template arguments
	type args struct {
		TicketTitle              string
		AssigneeActionId         string
		ReporterActionId         string
		CreateJiraTicketActionId string
		IssueTypeActionId        string
	}

	my_args := args{
		TicketTitle:              ticketTitle,
		AssigneeActionId:         AssigneeActionId,
		ReporterActionId:         ReporterActionId,
		CreateJiraTicketActionId: CreateJiraTicketActionId,
		IssueTypeActionId:        IssueTypeActionId,
	}

	tpl := renderTemplate(slashCommandAssets, "jiraAssets/createJiraIssue.json", my_args)

	// we convert the view into a message struct
	view := slack.Msg{}

	str, _ := ioutil.ReadAll(&tpl)
	_ = json.Unmarshal(str, &view)

	// We only return the block because of the way the PostEphemeral function works
	// we are going to use slack.MsgOptionBlocks in the controller
	return view.Blocks.BlockSet
}

package slacktool

import (
	"github.com/slack-go/slack"
)

const (
	SlackCommandCtx = "slackCommand"
)

var slackClient *slack.Client

func Init(botToken string, opts ...slack.Option) {
	slackClient = slack.New(botToken, opts...)
}

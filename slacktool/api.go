package slacktool

import (
	"github.com/blackironj/jislack/config"
	"github.com/slack-go/slack"
)

const (
	SlackCommandCtx = "slackCommand"
)

var slackClient *slack.Client

func InitSlackAPI(opts ...slack.Option) {
	slackClient = slack.New(config.Get().Slack.BotToken, opts...)
}

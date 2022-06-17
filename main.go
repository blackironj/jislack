package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack/socketmode"

	"github.com/blackironj/jislack/config"
	"github.com/blackironj/jislack/jiratool"
	"github.com/blackironj/jislack/slacktool/controllers"
	"github.com/blackironj/jislack/slacktool/drivers"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	config.InitCfg("config/config.yml")
	cfg := config.Get()

	jiratool.Init(cfg.Jira.User, cfg.Jira.ApiToken, cfg.Jira.BaseURL)
	client, err := drivers.ConnectToSlackViaSocketmode(cfg.Slack.AppToken, cfg.Slack.BotToken)
	if err != nil {
		log.Error().
			Str("error", err.Error()).
			Msg("Unable to connect to slack")

		os.Exit(1)
	}

	// Inject Deps in router
	socketmodeHandler := socketmode.NewSocketmodeHandler(client)

	// Build Slack Slash Command in Golang Using Socket Mode
	controllers.NewSlashCommandController(socketmodeHandler)

	_ = socketmodeHandler.RunEventLoop()
}

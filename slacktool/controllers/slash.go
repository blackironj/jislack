package controllers

import (
	"log"
	"strings"

	"github.com/blackironj/jislack/slacktool/views"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

// We create a sctucture to let us use dependency injection
type SlashCommandController struct {
	EventHandler *socketmode.SocketmodeHandler
}

func NewSlashCommandController(eventhandler *socketmode.SocketmodeHandler) SlashCommandController {
	// we need to cast our socketmode.Event into a SlashCommand
	c := SlashCommandController{
		EventHandler: eventhandler,
	}

	// Register callback for the command /jislack
	c.EventHandler.HandleSlashCommand(
		"/jislack",
		c.createJiraTicketView,
	)

	c.EventHandler.HandleInteractionBlockAction(
		views.CreateJiraTicketActionId,
		c.createJiraTicket,
	)

	return c
}

func (c SlashCommandController) createJiraTicketView(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into a Slash Command
	command, ok := evt.Data.(slack.SlashCommand)

	if ok != true {
		log.Printf("ERROR converting event to Slash Command: %v", ok)
	}

	title := strings.TrimSpace(command.Text)
	// Make sure to respond to the server to avoid an error
	clt.Ack(*evt.Request)

	// create the view using block-kit
	blocks := views.CreateJiraInfoView(title)

	// Post ephemeral message
	_, _, err := clt.Client.PostMessage(
		command.ChannelID,
		slack.MsgOptionBlocks(blocks...),
		slack.MsgOptionResponseURL(command.ResponseURL, slack.ResponseTypeEphemeral),
	)

	// Handle errors
	if err != nil {
		log.Printf("ERROR while sending message for /rocket: %v", err)
	}

}

func (c SlashCommandController) rejectTikcet()

func (c SlashCommandController) createJiraTicket(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into a Slash Command
	interaction := evt.Data.(slack.InteractionCallback)

	// Make sure to respond to the server to avoid an error
	clt.Ack(*evt.Request)

	for _, val := range interaction.BlockActionState.Values {
		for key, val2 := range val {
			log.Println(key)
			log.Println(val2.ActionID)
		}
	}
}

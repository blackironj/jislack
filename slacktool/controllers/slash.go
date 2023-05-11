package controllers

import (
	"log"

	"github.com/blackironj/jislack/slacktool/views"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

const A = "A test"

// We create a sctucture to let us use dependency injection
type SlashCommandController struct {
	EventHandler *socketmode.SocketmodeHandler
	B_CODE       string
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
		views.CancelJiraTicketActionId,
		c.cancelCreatingJiraTikcet,
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

	// Make sure to respond to the server to avoid an error
	clt.Ack(*evt.Request)

	title := command.Text
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
		log.Printf("ERROR while sending message for /jislack: %v", err)
	}
}

func (c SlashCommandController) cancelCreatingJiraTikcet(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into a Slash Command
	interaction := evt.Data.(slack.InteractionCallback)

	// Make sure to respond to the server to avoid an error
	clt.Ack(*evt.Request)

	blocks := views.CancelView()
	// Post ephemeral message
	_, _, err := clt.Client.PostMessage(
		interaction.Container.ChannelID,
		slack.MsgOptionBlocks(blocks...),
		slack.MsgOptionResponseURL(interaction.ResponseURL, slack.ResponseTypeEphemeral),
		slack.MsgOptionReplaceOriginal(interaction.ResponseURL),
	)

	// Handle errors
	if err != nil {
		log.Printf("ERROR while sending message for /jislack cancel view: %v", err)
	}
}

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

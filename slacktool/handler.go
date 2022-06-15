package slacktool

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/blackironj/jislack/config"
	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
)

const (
	_version = "0.0.1"
)

func ValidateSlackCommandMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.Get()
		r := c.Request

		verifier, err := slack.NewSecretsVerifier(r.Header, cfg.Slack.SigningSecret)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		r.Body = ioutil.NopCloser(io.TeeReader(r.Body, &verifier))
		s, err := slack.SlashCommandParse(r)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := verifier.Ensure(); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set(SlackCommandCtx, s)
		c.Next()
	}
}

func CommandHandler(c *gin.Context) {
	s := c.MustGet(SlackCommandCtx).(slack.SlashCommand)

	parsedCommand := strings.Split(s.Text, " ")
	op := parsedCommand[0]

	switch op {
	case "help":
	case "version":
		c.JSON(http.StatusOK, &slack.Msg{Text: "version: " + _version})
	default:
		c.JSON(http.StatusOK, &slack.Msg{Text: "invalid command"})
	}
}

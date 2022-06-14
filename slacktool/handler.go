package slacktool

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/blackironj/jislack/config"
	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
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

		if err := verifier.Ensure(); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		r.Body = ioutil.NopCloser(io.TeeReader(r.Body, &verifier))
		s, err := slack.SlashCommandParse(r)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.Set(SlackCommandCtx, s)
		c.Next()
	}
}

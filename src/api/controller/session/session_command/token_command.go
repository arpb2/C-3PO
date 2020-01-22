package session_command

import (
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"net/http"
)

type TokenCommand struct {
	Context      *http_wrapper.Context
	TokenHandler auth.TokenHandler
	InputStream  chan *model.AuthenticatedUser
	OutputStream chan string
}

func (c *TokenCommand) Name() string {
	return "create_token_command"
}

func (c *TokenCommand) Run() error {
	user := <-c.InputStream

	token, tokenErr := c.TokenHandler.Create(&auth.Token{
		UserId: user.Id,
	})

	if tokenErr != nil {
		if tokenErr.Status == http.StatusInternalServerError {
			return tokenErr.Error
		} else {
			controller.Halt(c.Context, tokenErr.Status, tokenErr.Error.Error())
			return nil
		}
	}

	c.OutputStream <- token
	return nil
}

func (c *TokenCommand) Fallback(err error) error {
	return err
}

func CreateTokenCommand(ctx *http_wrapper.Context,
						tokenHandler auth.TokenHandler,
						inputStream chan *model.AuthenticatedUser) *TokenCommand {
	return &TokenCommand{
		Context:      ctx,
		TokenHandler: tokenHandler,
		InputStream:  inputStream,
		OutputStream: make(chan string, 1),
	}
}
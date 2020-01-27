package session

import (
	"github.com/arpb2/C-3PO/api/auth"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
)

type createTokenCommand struct {
	context      *http.Context
	tokenHandler auth.TokenHandler
	inputStream  <-chan *model.AuthenticatedUser

	OutputStream chan string
}

func (c *createTokenCommand) Name() string {
	return "create_token_command"
}

func (c *createTokenCommand) Run() error {
	defer close(c.OutputStream)

	token, err := c.tokenHandler.Create(&auth.Token{
		UserId: (<-c.inputStream).Id,
	})

	if err != nil {
		return err
	}

	c.OutputStream <- token
	return nil
}

func CreateCreateTokenCommand(ctx *http.Context,
	tokenHandler auth.TokenHandler,
	inputStream <-chan *model.AuthenticatedUser) *createTokenCommand {
	return &createTokenCommand{
		context:      ctx,
		tokenHandler: tokenHandler,
		inputStream:  inputStream,
		OutputStream: make(chan string, 1),
	}
}

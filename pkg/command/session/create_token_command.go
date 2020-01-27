package session

import (
	"github.com/arpb2/C-3PO/api/auth"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
)

type createSessionCommand struct {
	context      *http.Context
	tokenHandler auth.TokenHandler
	inputStream  <-chan *model.AuthenticatedUser

	OutputStream chan *model.Session
}

func (c *createSessionCommand) Name() string {
	return "create_token_command"
}

func (c *createSessionCommand) Run() error {
	defer close(c.OutputStream)

	userId := (<-c.inputStream).Id
	token, err := c.tokenHandler.Create(&auth.Token{
		UserId: userId,
	})

	if err != nil {
		return err
	}

	c.OutputStream <- &model.Session{
		UserId: userId,
		Token:  token,
	}
	return nil
}

func CreateCreateSessionCommand(ctx *http.Context,
	tokenHandler auth.TokenHandler,
	inputStream <-chan *model.AuthenticatedUser) *createSessionCommand {
	return &createSessionCommand{
		context:      ctx,
		tokenHandler: tokenHandler,
		inputStream:  inputStream,
		OutputStream: make(chan *model.Session, 1),
	}
}

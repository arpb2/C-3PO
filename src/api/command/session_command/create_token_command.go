package session_command

import (
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
)

type createTokenCommand struct {
	context      *http_wrapper.Context
	tokenHandler auth.TokenHandler
	inputStream  <-chan *model.AuthenticatedUser

	OutputStream chan string

	user         *model.AuthenticatedUser
}

func (c *createTokenCommand) Name() string {
	return "create_token_command"
}

func (c *createTokenCommand) Prepare() bool {
	c.user = <-c.inputStream

	if c.user == nil {
		close(c.OutputStream)
	}

	return c.user != nil
}

func (c *createTokenCommand) Run() error {
	defer close(c.OutputStream)

	token, err := c.tokenHandler.Create(&auth.Token{
		UserId: c.user.Id,
	})

	if err != nil {
		return command.HaltClientHttpError(c.context, err)
	}

	c.OutputStream <- token
	return nil
}

func (c *createTokenCommand) Fallback(err error) error {
	return err
}

func CreateCreateTokenCommand(ctx *http_wrapper.Context,
						tokenHandler auth.TokenHandler,
						inputStream <-chan *model.AuthenticatedUser) *createTokenCommand {
	return &createTokenCommand{
		context:      ctx,
		tokenHandler: tokenHandler,
		inputStream:  inputStream,
		OutputStream: make(chan string, 1),
	}
}
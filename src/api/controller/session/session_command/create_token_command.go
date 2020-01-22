package session_command

import (
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
)

type CreateTokenCommand struct {
	Context      *http_wrapper.Context
	TokenHandler auth.TokenHandler
	InputStream  <-chan *model.AuthenticatedUser
	OutputStream chan string
}

func (c *CreateTokenCommand) Name() string {
	return "create_token_command"
}

func (c *CreateTokenCommand) Run() error {
	defer close(c.OutputStream)
	user, openChan := <-c.InputStream

	if !openChan {
		return nil
	}

	token, err := c.TokenHandler.Create(&auth.Token{
		UserId: user.Id,
	})

	if err != nil {
		return controller.HaltExternalError(c.Context, err)
	}

	c.OutputStream <- token
	return nil
}

func (c *CreateTokenCommand) Fallback(err error) error {
	return err
}

func CreateCreateTokenCommand(ctx *http_wrapper.Context,
						tokenHandler auth.TokenHandler,
						inputStream <-chan *model.AuthenticatedUser) *CreateTokenCommand {
	return &CreateTokenCommand{
		Context:      ctx,
		TokenHandler: tokenHandler,
		InputStream:  inputStream,
		OutputStream: make(chan string, 1),
	}
}
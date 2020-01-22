package session_command

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type AuthenticateCommand struct {
	Context      *http_wrapper.Context
	Service      service.CredentialService
	Stream       chan *model.AuthenticatedUser
}

func (c *AuthenticateCommand) Name() string {
	return "authenticate_command"
}

func (c *AuthenticateCommand) Run() error {
	user := <-c.Stream
	userId, err := c.Service.Retrieve(user.Email, user.Password)

	if err != nil {
		return err
	}

	user.Id = userId
	c.Stream <- user
	return nil
}

func (c *AuthenticateCommand) Fallback(err error) error {
	return err
}

func CreateAuthenticateCommand(ctx *http_wrapper.Context,
							   service service.CredentialService,
	                           inputStream chan *model.AuthenticatedUser) *AuthenticateCommand {
	return &AuthenticateCommand{
		Context:      ctx,
		Service:      service,
		Stream:       inputStream,
	}
}
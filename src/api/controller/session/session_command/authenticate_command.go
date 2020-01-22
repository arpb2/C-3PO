package session_command

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type AuthenticateCommand struct {
	Context      *http_wrapper.Context
	Service      service.CredentialService
	InputStream  chan *model.AuthenticatedUser
	OutputStream  chan *model.AuthenticatedUser
}

func (c *AuthenticateCommand) Name() string {
	return "authenticate_command"
}

func (c *AuthenticateCommand) Run() error {
	defer close(c.OutputStream)
	user := <-c.InputStream
	userId, err := c.Service.Retrieve(user.Email, user.Password)

	if err != nil {
		return controller.HaltExternalError(c.Context, err)
	}

	user.Id = userId
	c.OutputStream <- user
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
		InputStream:  inputStream,
		OutputStream: make(chan *model.AuthenticatedUser, 1),
	}
}
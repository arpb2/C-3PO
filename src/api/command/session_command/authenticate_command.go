package session_command

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type authenticateCommand struct {
	context      *http_wrapper.Context
	service      service.CredentialService
	inputStream  chan *model.AuthenticatedUser

	OutputStream chan *model.AuthenticatedUser
}

func (c *authenticateCommand) Name() string {
	return "authenticate_command"
}

func (c *authenticateCommand) Run() error {
	defer close(c.OutputStream)

	user := <-c.inputStream

	userId, err := c.service.Retrieve(user.Email, user.Password)

	if err != nil {
		return err
	}

	user.Id = userId
	c.OutputStream <- user
	return nil
}

func CreateAuthenticateCommand(ctx *http_wrapper.Context,
							   service service.CredentialService,
	                           inputStream chan *model.AuthenticatedUser) *authenticateCommand {
	return &authenticateCommand{
		context:      ctx,
		service:      service,
		inputStream:  inputStream,
		OutputStream: make(chan *model.AuthenticatedUser, 1),
	}
}
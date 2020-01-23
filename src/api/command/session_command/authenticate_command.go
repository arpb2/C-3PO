package session_command

import (
	"github.com/arpb2/C-3PO/src/api/command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type authenticateCommand struct {
	context      *http_wrapper.Context
	service      service.CredentialService
	inputStream  chan *model.AuthenticatedUser

	OutputStream chan *model.AuthenticatedUser

	user         *model.AuthenticatedUser
}

func (c *authenticateCommand) Name() string {
	return "authenticate_command"
}

func (c *authenticateCommand) Prepare() bool {
	c.user = <-c.inputStream

	if c.user == nil {
		close(c.OutputStream)
	}

	return c.user != nil
}

func (c *authenticateCommand) Run() error {
	defer close(c.OutputStream)

	userId, err := c.service.Retrieve(c.user.Email, c.user.Password)

	if err != nil {
		return command.HaltExternalError(c.context, err)
	}

	c.user.Id = userId
	c.OutputStream <- c.user
	return nil
}

func (c *authenticateCommand) Fallback(err error) error {
	return err
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
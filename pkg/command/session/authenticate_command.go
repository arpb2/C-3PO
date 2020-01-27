package session

import (
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	credentialservice "github.com/arpb2/C-3PO/api/service/credential"
)

type authenticateCommand struct {
	context     *http.Context
	service     credentialservice.Service
	inputStream chan *model.AuthenticatedUser

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

func CreateAuthenticateCommand(ctx *http.Context,
	service credentialservice.Service,
	inputStream chan *model.AuthenticatedUser) *authenticateCommand {
	return &authenticateCommand{
		context:      ctx,
		service:      service,
		inputStream:  inputStream,
		OutputStream: make(chan *model.AuthenticatedUser, 1),
	}
}

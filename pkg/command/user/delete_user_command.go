package user

import (
	"github.com/arpb2/C-3PO/api/http"
	userservice "github.com/arpb2/C-3PO/api/service/user"
)

type deleteUserCommand struct {
	context     *http.Context
	service     userservice.Service
	inputStream <-chan uint

	OutputStream chan bool
}

func (c *deleteUserCommand) Name() string {
	return "delete_user_command"
}

func (c *deleteUserCommand) Run() error {
	defer close(c.OutputStream)

	err := c.service.DeleteUser(<-c.inputStream)

	if err != nil {
		return err
	}

	c.OutputStream <- true
	return nil
}

func CreateDeleteUserCommand(ctx *http.Context,
	service userservice.Service,
	inputStream <-chan uint) *deleteUserCommand {
	return &deleteUserCommand{
		context:      ctx,
		service:      service,
		inputStream:  inputStream,
		OutputStream: make(chan bool, 1),
	}
}

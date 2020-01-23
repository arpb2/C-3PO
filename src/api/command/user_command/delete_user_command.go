package user_command

import (
	"github.com/arpb2/C-3PO/src/api/command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
)

type deleteUserCommand struct {
	context      *http_wrapper.Context
	service      service.UserService
	inputStream  <-chan uint

	OutputStream chan bool

	userId       uint
}

func (c *deleteUserCommand) Name() string {
	return "delete_user_command"
}

func (c *deleteUserCommand) Prepare() bool {
	userId, openChan := <-c.inputStream

	if !openChan {
		close(c.OutputStream)
		return false
	}

	c.userId = userId
	return true
}

func (c *deleteUserCommand) Run() error {
	defer close(c.OutputStream)

	err := c.service.DeleteUser(c.userId)

	if err != nil {
		return command.HaltClientHttpError(c.context, err)
	}

	c.OutputStream <- true
	return nil
}

func (c *deleteUserCommand) Fallback(err error) error {
	return err
}

func CreateDeleteUserCommand(ctx *http_wrapper.Context,
							 service service.UserService,
							 inputStream <-chan uint) *deleteUserCommand {
	return &deleteUserCommand{
		context:      ctx,
		service:      service,
		inputStream:  inputStream,
		OutputStream: make(chan bool, 1),
	}
}
package user_command

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
)

type DeleteUserCommand struct {
	Context      *http_wrapper.Context
	Service      service.UserService
	InputStream  <-chan uint
}

func (c *DeleteUserCommand) Name() string {
	return "delete_user_command"
}

func (c *DeleteUserCommand) Run() error {
	err := c.Service.DeleteUser(<-c.InputStream)

	if err != nil {
		return err
	}

	return nil
}

func (c *DeleteUserCommand) Fallback(err error) error {
	return err
}

func CreateDeleteUserCommand(ctx *http_wrapper.Context,
							 service service.UserService,
							 inputStream <-chan uint) *DeleteUserCommand {
	return &DeleteUserCommand{
		Context:      ctx,
		Service:      service,
		InputStream:  inputStream,
	}
}
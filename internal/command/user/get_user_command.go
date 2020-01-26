package user_command

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/api/service/user"
)

type getUserCommand struct {
	context     *http_wrapper.Context
	service     user_service.Service
	inputStream <-chan uint

	OutputStream chan *model.User
}

func (c *getUserCommand) Name() string {
	return "get_user_command"
}

func (c *getUserCommand) Run() error {
	defer close(c.OutputStream)

	user, err := c.service.GetUser(<-c.inputStream)

	if err != nil {
		return err
	}

	if user == nil {
		return http_wrapper.CreateNotFoundError()
	}

	c.OutputStream <- user
	return nil
}

func CreateGetUserCommand(ctx *http_wrapper.Context,
	service user_service.Service,
	inputStream <-chan uint) *getUserCommand {
	return &getUserCommand{
		context:      ctx,
		service:      service,
		inputStream:  inputStream,
		OutputStream: make(chan *model.User, 1),
	}
}

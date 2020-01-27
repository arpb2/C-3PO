package user

import (
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	userservice "github.com/arpb2/C-3PO/api/service/user"
)

type getUserCommand struct {
	context     *http.Context
	service     userservice.Service
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
		return http.CreateNotFoundError()
	}

	c.OutputStream <- user
	return nil
}

func CreateGetUserCommand(ctx *http.Context,
	service userservice.Service,
	inputStream <-chan uint) *getUserCommand {
	return &getUserCommand{
		context:      ctx,
		service:      service,
		inputStream:  inputStream,
		OutputStream: make(chan *model.User, 1),
	}
}

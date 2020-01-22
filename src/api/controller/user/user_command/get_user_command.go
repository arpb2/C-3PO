package user_command

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type GetUserCommand struct {
	Context      *http_wrapper.Context
	Service      service.UserService
	InputStream  <-chan uint
	OutputStream chan *model.User
}

func (c *GetUserCommand) Name() string {
	return "get_user_command"
}

func (c *GetUserCommand) Run() error {
	defer close(c.OutputStream)
	userId, openChan := <-c.InputStream

	if !openChan {
		return nil
	}

	user, err := c.Service.GetUser(userId)

	if err != nil {
		return controller.HaltExternalError(c.Context, err)
	}

	if user == nil {
		return controller.HaltExternalError(c.Context, http_wrapper.CreateNotFoundError())
	}

	c.OutputStream <- user
	return nil
}

func (c *GetUserCommand) Fallback(err error) error {
	return err
}

func CreateGetUserCommand(ctx *http_wrapper.Context,
						  service service.UserService,
						  inputStream <-chan uint) *GetUserCommand {
	return &GetUserCommand{
		Context:      ctx,
		Service:      service,
		InputStream:  inputStream,
		OutputStream: make(chan *model.User, 1),
	}
}
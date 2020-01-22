package user_command

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type CreateUserCommand struct {
	Context      *http_wrapper.Context
	Service      service.UserService
	InputStream  <-chan *model.AuthenticatedUser
	OutputStream chan *model.User
}

func (c *CreateUserCommand) Name() string {
	return "create_user_command"
}

func (c *CreateUserCommand) Run() error {
	defer close(c.OutputStream)
	authenticatedUser, openChan := <-c.InputStream

	if !openChan {
		return nil
	}

	user, err := c.Service.CreateUser(authenticatedUser)

	if err != nil || user == nil {
		return controller.HaltExternalError(c.Context, err)
	}

	c.OutputStream <- user
	return nil
}

func (c *CreateUserCommand) Fallback(err error) error {
	return err
}

func CreateCreateUserCommand(ctx *http_wrapper.Context,
							 service service.UserService,
							 inputStream <-chan *model.AuthenticatedUser) *CreateUserCommand {
	return &CreateUserCommand{
		Context:      ctx,
		Service:      service,
		InputStream:  inputStream,
		OutputStream: make(chan *model.User, 1),
	}
}
package user_command

import (
	"github.com/arpb2/C-3PO/src/api/command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type createUserCommand struct {
	context      *http_wrapper.Context
	service      service.UserService
	inputStream  <-chan *model.AuthenticatedUser

	OutputStream chan *model.User

	user         *model.AuthenticatedUser
}

func (c *createUserCommand) Name() string {
	return "create_user_command"
}

func (c *createUserCommand) Prepare() bool {
	c.user = <-c.inputStream

	if c.user == nil {
		close(c.OutputStream)
		return false
	}

	return true
}

func (c *createUserCommand) Run() error {
	defer close(c.OutputStream)

	user, err := c.service.CreateUser(c.user)

	if err != nil {
		return command.HaltExternalError(c.context, err)
	}

	if user == nil {
		return command.HaltExternalError(c.context, http_wrapper.CreateInternalError())
	}

	c.OutputStream <- user
	return nil
}

func (c *createUserCommand) Fallback(err error) error {
	return err
}

func CreateCreateUserCommand(ctx *http_wrapper.Context,
							 service service.UserService,
							 inputStream <-chan *model.AuthenticatedUser) *createUserCommand {
	return &createUserCommand{
		context:      ctx,
		service:      service,
		inputStream:  inputStream,
		OutputStream: make(chan *model.User, 1),
	}
}
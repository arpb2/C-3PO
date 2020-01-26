package user_command

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/api/service/user"
)

type updateUserCommand struct {
	context           *http_wrapper.Context
	service           user_service.Service
	userIdInputStream <-chan uint
	userInputStream   <-chan *model.AuthenticatedUser

	OutputStream chan *model.User
}

func (c *updateUserCommand) Name() string {
	return "update_user_command"
}

func (c *updateUserCommand) Run() error {
	defer close(c.OutputStream)

	authenticatedUser := <-c.userInputStream
	authenticatedUser.Id = <-c.userIdInputStream

	user, err := c.service.UpdateUser(authenticatedUser)

	if err != nil {
		return err
	}

	if user == nil {
		return http_wrapper.CreateInternalError()
	}

	c.OutputStream <- user
	return nil
}

func CreateUpdateUserCommand(ctx *http_wrapper.Context,
	service user_service.Service,
	userIdInputStream <-chan uint,
	userInputStream <-chan *model.AuthenticatedUser) *updateUserCommand {
	return &updateUserCommand{
		context:           ctx,
		service:           service,
		userIdInputStream: userIdInputStream,
		userInputStream:   userInputStream,
		OutputStream:      make(chan *model.User, 1),
	}
}

package user_command

import (
	"github.com/arpb2/C-3PO/src/api/command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type updateUserCommand struct {
	context           *http_wrapper.Context
	service           service.UserService
	userIdInputStream <-chan uint
	userInputStream   <-chan *model.AuthenticatedUser

	OutputStream      chan *model.User

	authenticatedUser *model.AuthenticatedUser
}

func (c *updateUserCommand) Name() string {
	return "update_user_command"
}

func (c *updateUserCommand) Prepare() bool {
	userId, openUserIdChan := <-c.userIdInputStream
	user, openUserChan := <-c.userInputStream

	if !openUserIdChan || !openUserChan {
		close(c.OutputStream)
		return false
	}

	c.authenticatedUser = user
	c.authenticatedUser.Id = userId
	return true
}

func (c *updateUserCommand) Run() error {
	defer close(c.OutputStream)

	user, err := c.service.UpdateUser(c.authenticatedUser)

	if err != nil {
		return command.HaltExternalError(c.context, err)
	}

	if user == nil {
		return command.HaltExternalError(c.context, http_wrapper.CreateInternalError())
	}

	c.OutputStream <- user
	return nil
}

func (c *updateUserCommand) Fallback(err error) error {
	return err
}

func CreateUpdateUserCommand(ctx *http_wrapper.Context,
							 service service.UserService,
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
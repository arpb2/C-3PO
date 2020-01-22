package user_command


import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type UpdateUserCommand struct {
	Context           *http_wrapper.Context
	Service           service.UserService
	UserIdInputStream <-chan uint
	UserInputStream   <-chan *model.AuthenticatedUser
	OutputStream      chan *model.User
}

func (c *UpdateUserCommand) Name() string {
	return "update_user_command"
}

func (c *UpdateUserCommand) Run() error {
	userId := <-c.UserIdInputStream
	authenticatedUser := <-c.UserInputStream

	authenticatedUser.Id = userId

	user, err := c.Service.UpdateUser(authenticatedUser)

	if err != nil {
		return err
	}

	if user == nil {
		return http_wrapper.CreateInternalError()
	}

	c.OutputStream <- user
	return nil
}

func (c *UpdateUserCommand) Fallback(err error) error {
	return controller.HaltError(c.Context, err)
}

func CreateUpdateUserCommand(ctx *http_wrapper.Context,
							 service service.UserService,
							 userIdInputStream <-chan uint,
							 userInputStream <-chan *model.AuthenticatedUser) *UpdateUserCommand {
	return &UpdateUserCommand{
		Context:         ctx,
		Service:         service,
		UserIdInputStream: userIdInputStream,
		UserInputStream: userInputStream,
		OutputStream:    make(chan *model.User, 1),
	}
}
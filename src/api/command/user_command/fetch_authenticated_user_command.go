package user_command

import (
	"github.com/arpb2/C-3PO/src/api/command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
)

type fetchAuthenticatedUserCommand struct {
	context      *http_wrapper.Context

	OutputStream chan *model.AuthenticatedUser
}

func (c *fetchAuthenticatedUserCommand) Name() string {
	return "fetch_user_command"
}

func (c *fetchAuthenticatedUserCommand) Prepare() bool {
	return true
}

func (c *fetchAuthenticatedUserCommand) Run() error {
	defer close(c.OutputStream)
	var authenticatedUser model.AuthenticatedUser

	if err := c.context.ReadBody(&authenticatedUser); err != nil {
		return command.HaltExternalError(c.context, http_wrapper.CreateBadRequestError("malformed body"))
	}

	c.OutputStream <- &authenticatedUser
	return nil
}

func (c *fetchAuthenticatedUserCommand) Fallback(err error) error {
	return err
}

func CreateFetchAuthenticatedUserCommand(ctx *http_wrapper.Context) *fetchAuthenticatedUserCommand {
	return &fetchAuthenticatedUserCommand{
		context:      ctx,
		OutputStream: make(chan *model.AuthenticatedUser, 1),
	}
}
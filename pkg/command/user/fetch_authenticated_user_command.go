package user

import (
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
)

type fetchAuthenticatedUserCommand struct {
	context *http.Context

	OutputStream chan *model.AuthenticatedUser
}

func (c *fetchAuthenticatedUserCommand) Name() string {
	return "fetch_user_command"
}

func (c *fetchAuthenticatedUserCommand) Run() error {
	defer close(c.OutputStream)
	var authenticatedUser model.AuthenticatedUser

	if err := c.context.ReadBody(&authenticatedUser); err != nil {
		return http.CreateBadRequestError("malformed body")
	}

	c.OutputStream <- &authenticatedUser
	return nil
}

func CreateFetchAuthenticatedUserCommand(ctx *http.Context) *fetchAuthenticatedUserCommand {
	return &fetchAuthenticatedUserCommand{
		context:      ctx,
		OutputStream: make(chan *model.AuthenticatedUser, 1),
	}
}

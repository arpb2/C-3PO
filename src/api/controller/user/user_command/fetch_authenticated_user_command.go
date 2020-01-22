package user_command

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
)

type FetchAuthenticatedUserCommand struct {
	Context      *http_wrapper.Context
	OutputStream chan *model.AuthenticatedUser
}

func (c *FetchAuthenticatedUserCommand) Name() string {
	return "fetch_user_command"
}

func (c *FetchAuthenticatedUserCommand) Run() error {
	var authenticatedUser model.AuthenticatedUser

	if err := c.Context.ReadBody(&authenticatedUser); err != nil {
		return http_wrapper.CreateBadRequestError("malformed body")
	}

	c.OutputStream <- &authenticatedUser
	return nil
}

func (c *FetchAuthenticatedUserCommand) Fallback(err error) error {
	return controller.HaltError(c.Context, err)
}

func CreateFetchAuthenticatedUserCommand(ctx *http_wrapper.Context) *FetchAuthenticatedUserCommand {
	return &FetchAuthenticatedUserCommand{
		Context:      ctx,
		OutputStream: make(chan *model.AuthenticatedUser, 1),
	}
}
package session_command

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"net/http"
)

type FetchUserCommand struct {
	Context      *http_wrapper.Context
	OutputStream chan *model.AuthenticatedUser
}

func (c *FetchUserCommand) Name() string {
	return "fetch_user_command"
}

func (c *FetchUserCommand) Run() error {
	var authenticatedUser model.AuthenticatedUser

	if err := c.Context.ReadBody(&authenticatedUser); err != nil {
		controller.Halt(c.Context, http.StatusBadRequest, "malformed body")
		return nil
	}

	c.OutputStream <- &authenticatedUser
	return nil
}

func (c *FetchUserCommand) Fallback(err error) error {
	return err
}

func CreateFetchUserCommand(ctx *http_wrapper.Context) *FetchUserCommand {
	return &FetchUserCommand{
		Context:      ctx,
		OutputStream: make(chan *model.AuthenticatedUser, 1),
	}
}
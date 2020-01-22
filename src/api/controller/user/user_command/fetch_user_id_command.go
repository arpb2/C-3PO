package user_command

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"strconv"
)

type FetchUserIdCommand struct {
	Context      *http_wrapper.Context
	OutputStream chan uint
}

func (c *FetchUserIdCommand) Name() string {
	return "fetch_user_id_command"
}

func (c *FetchUserIdCommand) Run() error {
	defer close(c.OutputStream)
	userId := c.Context.GetParameter("user_id")

	if userId == "" {
		return controller.HaltExternalError(c.Context, http_wrapper.CreateBadRequestError("'user_id' empty"))
	}

	userIdUint, err := strconv.ParseUint(userId, 10, 64)

	if err != nil {
		return controller.HaltExternalError(c.Context, http_wrapper.CreateBadRequestError("'user_id' malformed, expecting a positive number"))
	}

	c.OutputStream <- uint(userIdUint)
	return nil
}

func (c *FetchUserIdCommand) Fallback(err error) error {
	return err
}

func CreateFetchUserIdCommand(ctx *http_wrapper.Context) *FetchUserIdCommand {
	return &FetchUserIdCommand{
		Context:      ctx,
		OutputStream: make(chan uint, 1),
	}
}
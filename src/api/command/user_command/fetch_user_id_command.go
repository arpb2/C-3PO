package user_command

import (
	"github.com/arpb2/C-3PO/src/api/command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"strconv"
)

type fetchUserIdCommand struct {
	context      *http_wrapper.Context
	OutputStream chan uint
}

func (c *fetchUserIdCommand) Name() string {
	return "fetch_user_id_command"
}

func (c *fetchUserIdCommand) Prepare() bool {
	return true
}

func (c *fetchUserIdCommand) Run() error {
	defer close(c.OutputStream)
	userId := c.context.GetParameter("user_id")

	if userId == "" {
		return command.HaltClientHttpError(c.context, http_wrapper.CreateBadRequestError("'user_id' empty"))
	}

	userIdUint, err := strconv.ParseUint(userId, 10, 64)

	if err != nil {
		return command.HaltClientHttpError(c.context, http_wrapper.CreateBadRequestError("'user_id' malformed, expecting a positive number"))
	}

	c.OutputStream <- uint(userIdUint)
	return nil
}

func (c *fetchUserIdCommand) Fallback(err error) error {
	return err
}

func CreateFetchUserIdCommand(ctx *http_wrapper.Context) *fetchUserIdCommand {
	return &fetchUserIdCommand{
		context:      ctx,
		OutputStream: make(chan uint, 1),
	}
}
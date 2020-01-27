package user

import (
	"strconv"

	"github.com/arpb2/C-3PO/api/http"
)

type fetchUserIdCommand struct {
	context      *http.Context
	OutputStream chan uint
}

func (c *fetchUserIdCommand) Name() string {
	return "fetch_user_id_command"
}

func (c *fetchUserIdCommand) Run() error {
	defer close(c.OutputStream)
	userId := c.context.GetParameter("user_id")

	if userId == "" {
		return http.CreateBadRequestError("'user_id' empty")
	}

	userIdUint, err := strconv.ParseUint(userId, 10, 64)

	if err != nil {
		return http.CreateBadRequestError("'user_id' malformed, expecting a positive number")
	}

	c.OutputStream <- uint(userIdUint)
	return nil
}

func CreateFetchUserIdCommand(ctx *http.Context) *fetchUserIdCommand {
	return &fetchUserIdCommand{
		context:      ctx,
		OutputStream: make(chan uint, 1),
	}
}

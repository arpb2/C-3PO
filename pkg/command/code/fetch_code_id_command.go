package code_command

import (
	"strconv"

	"github.com/arpb2/C-3PO/api/http_wrapper"
)

type fetchCodeIdCommand struct {
	context *http_wrapper.Context

	OutputStream chan uint
}

func (c *fetchCodeIdCommand) Name() string {
	return "fetch_code_id_command"
}

func (c *fetchCodeIdCommand) Run() error {
	defer close(c.OutputStream)
	codeId := c.context.GetParameter("code_id")

	if codeId == "" {
		return http_wrapper.CreateBadRequestError("'code_id' empty")
	}

	codeIdUint, err := strconv.ParseUint(codeId, 10, 64)

	if err != nil {
		return http_wrapper.CreateBadRequestError("'code_id' malformed, expecting a positive number")
	}

	c.OutputStream <- uint(codeIdUint)
	return nil
}

func CreateFetchCodeIdCommand(ctx *http_wrapper.Context) *fetchCodeIdCommand {
	return &fetchCodeIdCommand{
		context:      ctx,
		OutputStream: make(chan uint, 1),
	}
}

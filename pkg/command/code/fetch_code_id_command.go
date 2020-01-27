package code

import (
	"strconv"

	"github.com/arpb2/C-3PO/api/http"
)

type fetchCodeIdCommand struct {
	context *http.Context

	OutputStream chan uint
}

func (c *fetchCodeIdCommand) Name() string {
	return "fetch_code_id_command"
}

func (c *fetchCodeIdCommand) Run() error {
	defer close(c.OutputStream)
	codeId := c.context.GetParameter("code_id")

	if codeId == "" {
		return http.CreateBadRequestError("'code_id' empty")
	}

	codeIdUint, err := strconv.ParseUint(codeId, 10, 64)

	if err != nil {
		return http.CreateBadRequestError("'code_id' malformed, expecting a positive number")
	}

	c.OutputStream <- uint(codeIdUint)
	return nil
}

func CreateFetchCodeIdCommand(ctx *http.Context) *fetchCodeIdCommand {
	return &fetchCodeIdCommand{
		context:      ctx,
		OutputStream: make(chan uint, 1),
	}
}

package code_command

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"strconv"
)

type FetchCodeIdCommand struct {
	Context      *http_wrapper.Context
	OutputStream chan uint
}

func (c *FetchCodeIdCommand) Name() string {
	return "fetch_code_id_command"
}

func (c *FetchCodeIdCommand) Run() error {
	defer close(c.OutputStream)
	codeId := c.Context.GetParameter("code_id")

	if codeId == "" {
		return controller.HaltExternalError(c.Context, http_wrapper.CreateBadRequestError("'code_id' empty"))
	}

	codeIdUint, err := strconv.ParseUint(codeId, 10, 64)

	if err != nil {
		return controller.HaltExternalError(c.Context, http_wrapper.CreateBadRequestError("'code_id' malformed, expecting a positive number"))
	}

	c.OutputStream <- uint(codeIdUint)
	return nil
}

func (c *FetchCodeIdCommand) Fallback(err error) error {
	return err
}

func CreateFetchCodeIdCommand(ctx *http_wrapper.Context) *FetchCodeIdCommand {
	return &FetchCodeIdCommand{
		Context:      ctx,
		OutputStream: make(chan uint, 1),
	}
}
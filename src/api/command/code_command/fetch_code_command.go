package code_command

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
)

type fetchCodeCommand struct {
	context      *http_wrapper.Context

	OutputStream chan string
}

func (c *fetchCodeCommand) Name() string {
	return "fetch_code_command"
}

func (c *fetchCodeCommand) Run() error {
	defer close(c.OutputStream)
	code, exists := c.context.GetFormData("code")

	if !exists {
		return http_wrapper.CreateBadRequestError("'code' part not found")
	}

	c.OutputStream <- code
	return nil
}

func CreateFetchCodeCommand(ctx *http_wrapper.Context) *fetchCodeCommand {
	return &fetchCodeCommand{
		context:      ctx,
		OutputStream: make(chan string, 1),
	}
}
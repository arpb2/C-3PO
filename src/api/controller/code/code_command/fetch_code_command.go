package code_command

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
)

type FetchCodeCommand struct {
	Context      *http_wrapper.Context
	OutputStream chan string
}

func (c *FetchCodeCommand) Name() string {
	return "fetch_code_command"
}

func (c *FetchCodeCommand) Run() error {
	defer close(c.OutputStream)
	code, exists := c.Context.GetFormData("code")

	if !exists {
		return controller.HaltExternalError(c.Context, http_wrapper.CreateBadRequestError("'code' part not found"))
	}

	c.OutputStream <- code
	return nil
}

func (c *FetchCodeCommand) Fallback(err error) error {
	return err
}

func CreateFetchCodeCommand(ctx *http_wrapper.Context) *FetchCodeCommand {
	return &FetchCodeCommand{
		Context:      ctx,
		OutputStream: make(chan string, 1),
	}
}
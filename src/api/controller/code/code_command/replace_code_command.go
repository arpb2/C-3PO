package code_command

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type ReplaceCodeCommand struct {
	Context            *http_wrapper.Context
	Service            service.CodeService
	CodeInputStream    <-chan *model.Code
	OutputStream       chan *model.Code
}

func (c *ReplaceCodeCommand) Name() string {
	return "replace_code_command"
}

func (c *ReplaceCodeCommand) Run() error {
	defer close(c.OutputStream)
	code, openChan := <-c.CodeInputStream

	if !openChan {
		return nil
	}

	err := c.Service.ReplaceCode(code)

	if err != nil {
		return controller.HaltExternalError(c.Context, err)
	}

	c.OutputStream <- code
	return nil
}

func (c *ReplaceCodeCommand) Fallback(err error) error {
	return err
}

func CreateReplaceCodeCommand(ctx *http_wrapper.Context,
							  service service.CodeService,
							  codeInputStream <-chan *model.Code) *ReplaceCodeCommand {
	return &ReplaceCodeCommand{
		Context:           ctx,
		Service:           service,
		CodeInputStream:   codeInputStream,
		OutputStream:      make(chan *model.Code, 1),
	}
}
package code_command

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type CreateCodeCommand struct {
	Context            *http_wrapper.Context
	Service            service.CodeService
	UserIdInputStream  <-chan uint
	CodeInputStream    <-chan string
	OutputStream       chan *model.Code
}

func (c *CreateCodeCommand) Name() string {
	return "create_code_command"
}

func (c *CreateCodeCommand) Run() error {
	defer close(c.OutputStream)
	userId, openUserIdChan := <-c.UserIdInputStream
	rawCode, openRawCodeChan := <-c.CodeInputStream

	if !openRawCodeChan && !openUserIdChan {
		return nil
	}

	code, err := c.Service.CreateCode(userId, rawCode)

	if err != nil {
		return controller.HaltExternalError(c.Context, err)
	}

	c.OutputStream <- code
	return nil
}

func (c *CreateCodeCommand) Fallback(err error) error {
	return err
}

func CreateCreateCodeCommand(ctx *http_wrapper.Context,
							 service service.CodeService,
							 userIdInputStream <-chan uint,
							 codeInputStream <-chan string) *CreateCodeCommand {
	return &CreateCodeCommand{
		Context:           ctx,
		Service:           service,
		UserIdInputStream: userIdInputStream,
		CodeInputStream:   codeInputStream,
		OutputStream:      make(chan *model.Code, 1),
	}
}
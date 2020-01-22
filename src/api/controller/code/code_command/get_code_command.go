package code_command

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type GetCodeCommand struct {
	Context            *http_wrapper.Context
	Service            service.CodeService
	UserIdInputStream  <-chan uint
	CodeIdInputStream  <-chan uint
	OutputStream       chan *model.Code
}

func (c *GetCodeCommand) Name() string {
	return "get_code_command"
}

func (c *GetCodeCommand) Run() error {
	defer close(c.OutputStream)
	userId, openUserIdChan := <-c.UserIdInputStream
	codeId, openCodeIdChan := <-c.CodeIdInputStream

	if !openUserIdChan && !openCodeIdChan {
		return nil
	}

	code, err := c.Service.GetCode(userId, codeId)

	if err != nil {
		return controller.HaltExternalError(c.Context, err)
	}

	if code == nil {
		return controller.HaltExternalError(c.Context, http_wrapper.CreateNotFoundError())
	}

	c.OutputStream <- code
	return nil
}

func (c *GetCodeCommand) Fallback(err error) error {
	return err
}

func CreateGetCodeCommand(ctx *http_wrapper.Context,
						  service service.CodeService,
						  userIdInputStream <-chan uint,
						  codeIdInputStream <-chan uint) *GetCodeCommand {
	return &GetCodeCommand{
		Context:           ctx,
		Service:           service,
		UserIdInputStream: userIdInputStream,
		CodeIdInputStream: codeIdInputStream,
		OutputStream:      make(chan *model.Code, 1),
	}
}
package code_command

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type getCodeCommand struct {
	context           *http_wrapper.Context
	service           service.CodeService
	userIdInputStream <-chan uint
	codeIdInputStream <-chan uint

	OutputStream      chan *model.Code
}

func (c *getCodeCommand) Name() string {
	return "get_code_command"
}

func (c *getCodeCommand) Run() error {
	defer close(c.OutputStream)

	code, err := c.service.GetCode(<-c.userIdInputStream, <-c.codeIdInputStream)

	if err != nil {
		return err
	}

	if code == nil {
		return http_wrapper.CreateNotFoundError()
	}

	c.OutputStream <- code
	return nil
}

func CreateGetCodeCommand(ctx *http_wrapper.Context,
						  service service.CodeService,
						  userIdInputStream <-chan uint,
						  codeIdInputStream <-chan uint) *getCodeCommand {
	return &getCodeCommand{
		context:           ctx,
		service:           service,
		userIdInputStream: userIdInputStream,
		codeIdInputStream: codeIdInputStream,
		OutputStream:      make(chan *model.Code, 1),
	}
}
package code_command

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type replaceCodeCommand struct {
	context         *http_wrapper.Context
	service         service.CodeService
	codeIdInputStream <-chan uint
	userIdInputStream <-chan uint
	codeInputStream   <-chan string

	OutputStream    chan *model.Code
}

func (c *replaceCodeCommand) Name() string {
	return "replace_code_command"
}

func (c *replaceCodeCommand) Run() error {
	defer close(c.OutputStream)

	code := &model.Code{
		Id:     <-c.codeIdInputStream,
		UserId: <-c.userIdInputStream,
		Code:   <-c.codeInputStream,
	}

	err := c.service.ReplaceCode(code)

	if err != nil {
		return err
	}

	c.OutputStream <- code
	return nil
}

func CreateReplaceCodeCommand(ctx *http_wrapper.Context,
							  service service.CodeService,
							  codeIdInputStream <-chan uint,
							  userIdInputStream <-chan uint,
							  codeInputStream   <-chan string) *replaceCodeCommand {
	return &replaceCodeCommand{
		context:           ctx,
		service:           service,
		codeInputStream:   codeInputStream,
		userIdInputStream: userIdInputStream,
		codeIdInputStream: codeIdInputStream,
		OutputStream:      make(chan *model.Code, 1),
	}
}
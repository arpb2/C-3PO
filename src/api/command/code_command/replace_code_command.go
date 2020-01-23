package code_command

import (
	"github.com/arpb2/C-3PO/src/api/command"
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

	code            *model.Code
}

func (c *replaceCodeCommand) Name() string {
	return "replace_code_command"
}

func (c *replaceCodeCommand) Prepare() bool {
	userId, openUserIdChan := <-c.userIdInputStream
	codeId, openCodeIdChan := <-c.codeIdInputStream
	code, openCodeChan := <-c.codeInputStream

	if !openUserIdChan || !openCodeIdChan || !openCodeChan {
		close(c.OutputStream)
		return false
	}

	c.code = &model.Code{
		Id:     codeId,
		UserId: userId,
		Code:   code,
	}
	return true
}

func (c *replaceCodeCommand) Run() error {
	defer close(c.OutputStream)

	err := c.service.ReplaceCode(c.code)

	if err != nil {
		return command.HaltExternalError(c.context, err)
	}

	c.OutputStream <- c.code
	return nil
}

func (c *replaceCodeCommand) Fallback(err error) error {
	return err
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
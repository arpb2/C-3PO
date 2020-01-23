package code_command

import (
	"github.com/arpb2/C-3PO/src/api/command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type createCodeCommand struct {
	context           *http_wrapper.Context
	service           service.CodeService
	userIdInputStream <-chan uint
	codeInputStream   <-chan string

	OutputStream      chan *model.Code

	userId            uint
	code              string
}

func (c *createCodeCommand) Name() string {
	return "create_code_command"
}

func (c *createCodeCommand) Prepare() bool {
	userId, openUserIdChan := <-c.userIdInputStream
	rawCode, openRawCodeChan := <-c.codeInputStream

	if !openRawCodeChan || !openUserIdChan {
		close(c.OutputStream)
		return false
	}

	c.userId = userId
	c.code = rawCode
	return true
}

func (c *createCodeCommand) Run() error {
	defer close(c.OutputStream)

	code, err := c.service.CreateCode(c.userId, c.code)

	if err != nil {
		return command.HaltExternalError(c.context, err)
	}

	c.OutputStream <- code
	return nil
}

func (c *createCodeCommand) Fallback(err error) error {
	return err
}

func CreateCreateCodeCommand(ctx *http_wrapper.Context,
							 service service.CodeService,
							 userIdInputStream <-chan uint,
							 codeInputStream <-chan string) *createCodeCommand {
	return &createCodeCommand{
		context:           ctx,
		service:           service,
		userIdInputStream: userIdInputStream,
		codeInputStream:   codeInputStream,
		OutputStream:      make(chan *model.Code, 1),
	}
}
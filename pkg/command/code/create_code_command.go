package code_command

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/model"
	code_service "github.com/arpb2/C-3PO/api/service/code"
)

type createCodeCommand struct {
	context           *http_wrapper.Context
	service           code_service.Service
	userIdInputStream <-chan uint
	codeInputStream   <-chan string

	OutputStream chan *model.Code
}

func (c *createCodeCommand) Name() string {
	return "create_code_command"
}

func (c *createCodeCommand) Run() error {
	defer close(c.OutputStream)

	code, err := c.service.CreateCode(<-c.userIdInputStream, <-c.codeInputStream)

	if err != nil {
		return err
	}

	c.OutputStream <- code
	return nil
}

func CreateCreateCodeCommand(ctx *http_wrapper.Context,
	service code_service.Service,
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

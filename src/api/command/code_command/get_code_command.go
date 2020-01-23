package code_command

import (
	"github.com/arpb2/C-3PO/src/api/command"
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

	userId            uint
	codeId            uint
}

func (c *getCodeCommand) Name() string {
	return "get_code_command"
}

func (c *getCodeCommand) Prepare() bool {
	userId, openUserIdChan := <-c.userIdInputStream
	codeId, openCodeIdChan := <-c.codeIdInputStream

	if !openCodeIdChan || !openUserIdChan {
		close(c.OutputStream)
		return false
	}

	c.userId = userId
	c.codeId = codeId
	return true
}

func (c *getCodeCommand) Run() error {
	defer close(c.OutputStream)

	code, err := c.service.GetCode(c.userId, c.codeId)

	if err != nil {
		return command.HaltClientHttpError(c.context, err)
	}

	if code == nil {
		return command.HaltClientHttpError(c.context, http_wrapper.CreateNotFoundError())
	}

	c.OutputStream <- code
	return nil
}

func (c *getCodeCommand) Fallback(err error) error {
	return err
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
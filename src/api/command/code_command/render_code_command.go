package code_command

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"net/http"
)

type renderCodeCommand struct {
	writer      http_wrapper.Writer
	inputStream chan *model.Code

	code        *model.Code
}

func (c *renderCodeCommand) Name() string {
	return "render_code_command"
}

func (c *renderCodeCommand) Prepare() bool {
	c.code = <-c.inputStream
	return c.code != nil
}

func (c *renderCodeCommand) Run() error {
	c.writer.WriteJson(http.StatusOK, *c.code)
	return nil
}

func (c *renderCodeCommand) Fallback(err error) error {
	return err
}

func CreateRenderCodeCommand(writer http_wrapper.Writer, inputStream chan *model.Code) *renderCodeCommand {
	return &renderCodeCommand{
		writer:      writer,
		inputStream: inputStream,
	}
}
package code_command

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/model"
	"net/http"
)

type renderCodeCommand struct {
	writer      http_wrapper.Writer
	inputStream chan *model.Code
}

func (c *renderCodeCommand) Name() string {
	return "render_code_command"
}

func (c *renderCodeCommand) Run() error {
	c.writer.WriteJson(http.StatusOK, *<-c.inputStream)
	return nil
}

func CreateRenderCodeCommand(writer http_wrapper.Writer, inputStream chan *model.Code) *renderCodeCommand {
	return &renderCodeCommand{
		writer:      writer,
		inputStream: inputStream,
	}
}

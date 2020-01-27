package code

import (
	"net/http"

	httpwrapper "github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
)

type renderCodeCommand struct {
	writer      httpwrapper.Writer
	inputStream chan *model.Code
}

func (c *renderCodeCommand) Name() string {
	return "render_code_command"
}

func (c *renderCodeCommand) Run() error {
	c.writer.WriteJson(http.StatusOK, *<-c.inputStream)
	return nil
}

func CreateRenderCodeCommand(writer httpwrapper.Writer, inputStream chan *model.Code) *renderCodeCommand {
	return &renderCodeCommand{
		writer:      writer,
		inputStream: inputStream,
	}
}

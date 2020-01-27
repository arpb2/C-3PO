package session

import (
	"net/http"

	"github.com/arpb2/C-3PO/api/model"

	httpwrapper "github.com/arpb2/C-3PO/api/http"
)

type renderSessionCommand struct {
	writer      httpwrapper.Writer
	inputStream <-chan *model.Session
}

func (c *renderSessionCommand) Name() string {
	return "render_session_command"
}

func (c *renderSessionCommand) Run() error {
	c.writer.WriteJson(http.StatusOK, *<-c.inputStream)
	return nil
}

func CreateRenderSessionCommand(writer httpwrapper.Writer, inputStream chan *model.Session) *renderSessionCommand {
	return &renderSessionCommand{
		writer:      writer,
		inputStream: inputStream,
	}
}

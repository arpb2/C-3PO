package session

import (
	"net/http"

	httpwrapper "github.com/arpb2/C-3PO/api/http"
)

type renderSessionCommand struct {
	writer      httpwrapper.Writer
	inputStream <-chan string
}

func (c *renderSessionCommand) Name() string {
	return "render_session_command"
}

func (c *renderSessionCommand) Run() error {
	c.writer.WriteJson(http.StatusOK, httpwrapper.Json{
		"token": <-c.inputStream,
	})
	return nil
}

func CreateRenderSessionCommand(writer httpwrapper.Writer, inputStream chan string) *renderSessionCommand {
	return &renderSessionCommand{
		writer:      writer,
		inputStream: inputStream,
	}
}

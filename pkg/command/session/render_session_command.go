package session_command

import (
	"net/http"

	"github.com/arpb2/C-3PO/api/http_wrapper"
)

type renderSessionCommand struct {
	writer      http_wrapper.Writer
	inputStream <-chan string
}

func (c *renderSessionCommand) Name() string {
	return "render_session_command"
}

func (c *renderSessionCommand) Run() error {
	c.writer.WriteJson(http.StatusOK, http_wrapper.Json{
		"token": <-c.inputStream,
	})
	return nil
}

func CreateRenderSessionCommand(writer http_wrapper.Writer,
	inputStream chan string) *renderSessionCommand {
	return &renderSessionCommand{
		writer:      writer,
		inputStream: inputStream,
	}
}

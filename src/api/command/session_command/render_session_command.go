package session_command

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"net/http"
)

type renderSessionCommand struct {
	writer      http_wrapper.Writer
	inputStream <-chan string

	data        string
}

func (c *renderSessionCommand) Name() string {
	return "render_session_command"
}

func (c *renderSessionCommand) Prepare() bool {
	c.data = <-c.inputStream
	return len(c.data) != 0
}

func (c *renderSessionCommand) Run() error {
	c.writer.WriteJson(http.StatusOK, http_wrapper.Json{
		"token": c.data,
	})
	return nil
}

func (c *renderSessionCommand) Fallback(err error) error {
	return err
}

func CreateRenderSessionCommand(writer http_wrapper.Writer,
	                            inputStream chan string) *renderSessionCommand {
	return &renderSessionCommand{
		writer:      writer,
		inputStream: inputStream,
	}
}
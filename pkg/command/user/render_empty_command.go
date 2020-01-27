package user_command

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"net/http"
)

type renderEmptyCommand struct {
	writer      http_wrapper.Writer
	inputStream <-chan bool
}

func (c *renderEmptyCommand) Name() string {
	return "render_empty_command"
}

func (c *renderEmptyCommand) Run() error {
	c.writer.WriteStatus(http.StatusOK)
	return nil
}

func CreateRenderEmptyCommand(writer http_wrapper.Writer,
	inputStream <-chan bool) *renderEmptyCommand {
	return &renderEmptyCommand{
		writer:      writer,
		inputStream: inputStream,
	}
}

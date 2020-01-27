package user

import (
	"net/http"

	httpwrapper "github.com/arpb2/C-3PO/api/http"
)

type renderEmptyCommand struct {
	writer      httpwrapper.Writer
	inputStream <-chan bool
}

func (c *renderEmptyCommand) Name() string {
	return "render_empty_command"
}

func (c *renderEmptyCommand) Run() error {
	c.writer.WriteStatus(http.StatusOK)
	return nil
}

func CreateRenderEmptyCommand(writer httpwrapper.Writer, inputStream <-chan bool) *renderEmptyCommand {
	return &renderEmptyCommand{
		writer:      writer,
		inputStream: inputStream,
	}
}

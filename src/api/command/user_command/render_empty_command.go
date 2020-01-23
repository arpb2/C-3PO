package user_command

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"net/http"
)

type renderEmptyCommand struct {
	writer      http_wrapper.Writer
	inputStream <-chan bool
}

func (c *renderEmptyCommand) Name() string {
	return "render_empty_command"
}

func (c *renderEmptyCommand) Prepare() bool {
	return <-c.inputStream
}

func (c *renderEmptyCommand) Run() error {
	c.writer.WriteStatus(http.StatusOK)
	return nil
}

func (c *renderEmptyCommand) Fallback(err error) error {
	return err
}

func CreateRenderEmptyCommand(writer http_wrapper.Writer,
							 inputStream <-chan bool) *renderEmptyCommand {
	return &renderEmptyCommand{
		writer:      writer,
		inputStream: inputStream,
	}
}
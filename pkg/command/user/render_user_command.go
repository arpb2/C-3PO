package user_command

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/model"
	"net/http"
)

type renderUserCommand struct {
	writer      http_wrapper.Writer
	inputStream <-chan *model.User
}

func (c *renderUserCommand) Name() string {
	return "render_user_command"
}

func (c *renderUserCommand) Run() error {
	c.writer.WriteJson(http.StatusOK, *<-c.inputStream)
	return nil
}

func CreateRenderUserCommand(writer http_wrapper.Writer,
	inputStream <-chan *model.User) *renderUserCommand {
	return &renderUserCommand{
		writer:      writer,
		inputStream: inputStream,
	}
}

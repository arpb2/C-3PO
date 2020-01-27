package user

import (
	"net/http"

	httpwrapper "github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
)

type renderUserCommand struct {
	writer      httpwrapper.Writer
	inputStream <-chan *model.User
}

func (c *renderUserCommand) Name() string {
	return "render_user_command"
}

func (c *renderUserCommand) Run() error {
	c.writer.WriteJson(http.StatusOK, *<-c.inputStream)
	return nil
}

func CreateRenderUserCommand(writer httpwrapper.Writer, inputStream <-chan *model.User) *renderUserCommand {
	return &renderUserCommand{
		writer:      writer,
		inputStream: inputStream,
	}
}

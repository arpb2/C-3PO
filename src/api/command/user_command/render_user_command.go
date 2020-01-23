package user_command

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"net/http"
)

type renderUserCommand struct {
	writer      http_wrapper.Writer
	inputStream <-chan *model.User

	user        *model.User
}

func (c *renderUserCommand) Name() string {
	return "render_user_command"
}

func (c *renderUserCommand) Prepare() bool {
	c.user = <-c.inputStream
	return c.user != nil
}

func (c *renderUserCommand) Run() error {
	c.writer.WriteJson(http.StatusOK, *c.user)
	return nil
}

func (c *renderUserCommand) Fallback(err error) error {
	return err
}

func CreateRenderUserCommand(writer http_wrapper.Writer,
							 inputStream <-chan *model.User) *renderUserCommand {
	return &renderUserCommand{
		writer:      writer,
		inputStream: inputStream,
	}
}
package user

import (
	"net/http"

	"github.com/arpb2/C-3PO/pkg/command"
	"github.com/saantiaguilera/go-pipeline"

	httpwrapper "github.com/arpb2/C-3PO/api/http"
)

type renderUserCommand struct{}

func (c *renderUserCommand) Name() string {
	return "render_user_command"
}

func (c *renderUserCommand) Run(ctx pipeline.Context) error {
	httpWriter, existsWriter := ctx.Get(command.TagHttpWriter)
	user, existsUser := ctx.Get(TagUser)

	if !existsWriter || !existsUser {
		return httpwrapper.CreateInternalError()
	}

	httpWriter.(httpwrapper.Writer).WriteJson(http.StatusOK, user)
	return nil
}

func CreateRenderUserCommand() pipeline.Step {
	return &renderUserCommand{}
}

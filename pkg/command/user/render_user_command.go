package user

import (
	"net/http"

	httppipeline "github.com/arpb2/C-3PO/pkg/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	httpwrapper "github.com/arpb2/C-3PO/api/http"
)

type renderUserCommand struct{}

func (c *renderUserCommand) Name() string {
	return "render_user_command"
}

func (c *renderUserCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpWriter, errWriter := ctxAware.GetWriter()
	user, errUser := ctxAware.GetUser(TagUser)

	if errWriter != nil || errUser != nil {
		return httpwrapper.CreateInternalError()
	}

	httpWriter.WriteJson(http.StatusOK, user)
	return nil
}

func CreateRenderUserCommand() pipeline.Step {
	return &renderUserCommand{}
}

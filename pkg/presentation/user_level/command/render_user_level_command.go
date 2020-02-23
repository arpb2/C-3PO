package command

import (
	"net/http"

	httppipeline "github.com/arpb2/C-3PO/pkg/infra/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
)

type renderUserLevelCommand struct{}

func (c *renderUserLevelCommand) Name() string {
	return "render_user_level_command"
}

func (c *renderUserLevelCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpWriter, errWriter := ctxAware.GetWriter()
	userLevel, errCode := ctxAware.GetUserLevel(TagUserLevel)

	if errWriter != nil || errCode != nil {
		return httpwrapper.CreateInternalError()
	}

	httpWriter.WriteJson(http.StatusOK, userLevel)
	return nil
}

func CreateRenderUserLevelCommand() pipeline.Step {
	return &renderUserLevelCommand{}
}

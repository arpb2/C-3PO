package command

import (
	"net/http"

	httppipeline "github.com/arpb2/C-3PO/pkg/infra/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
)

type renderLevelCommand struct{}

func (c *renderLevelCommand) Name() string {
	return "render_level_command"
}

func (c *renderLevelCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpWriter, errWriter := ctxAware.GetWriter()
	lvl, errLvl := ctxAware.GetLevel(TagLevel)

	if errWriter != nil || errLvl != nil {
		return httpwrapper.CreateInternalError()
	}

	httpWriter.WriteJson(http.StatusOK, lvl)
	return nil
}

func CreateRenderLevelCommand() pipeline.Step {
	return &renderLevelCommand{}
}

package level

import (
	"net/http"

	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
)

type renderLevelUseCase struct{}

func (c *renderLevelUseCase) Name() string {
	return "render_level_usecase"
}

func (c *renderLevelUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpWriter, errWriter := ctxAware.GetWriter()
	lvl, errLvl := ctxAware.GetLevel(TagLevel)

	if errWriter != nil || errLvl != nil {
		return httpwrapper.CreateInternalError()
	}

	httpWriter.WriteJson(http.StatusOK, lvl)
	return nil
}

func CreateRenderLevelUseCase() pipeline.Step {
	return &renderLevelUseCase{}
}

package user

import (
	"net/http"

	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
)

type renderUserLevelsUseCase struct{}

func (c *renderUserLevelsUseCase) Name() string {
	return "render_user_levels_usecase"
}

func (c *renderUserLevelsUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpWriter, errWriter := ctxAware.GetWriter()
	userLevels, errCode := ctxAware.GetUserLevels(TagUserLevels)

	if errWriter != nil || errCode != nil {
		return httpwrapper.CreateInternalError()
	}

	if userLevels == nil || len(userLevels) == 0 {
		httpWriter.WriteStatus(http.StatusNoContent)
	} else {
		httpWriter.WriteJson(http.StatusOK, userLevels)
	}
	return nil
}

func CreateRenderUserLevelsUseCase() pipeline.Step {
	return &renderUserLevelsUseCase{}
}

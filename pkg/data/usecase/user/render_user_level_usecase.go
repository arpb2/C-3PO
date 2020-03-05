package user

import (
	"net/http"

	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
)

type renderUserLevelUseCase struct{}

func (c *renderUserLevelUseCase) Name() string {
	return "render_user_level_usecase"
}

func (c *renderUserLevelUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpWriter, errWriter := ctxAware.GetWriter()
	userLevel, errCode := ctxAware.GetUserLevel(TagUserLevel)

	if errWriter != nil || errCode != nil {
		return httpwrapper.CreateInternalError()
	}

	httpWriter.WriteJson(http.StatusOK, userLevel)
	return nil
}

func CreateRenderUserLevelUseCase() pipeline.Step {
	return &renderUserLevelUseCase{}
}

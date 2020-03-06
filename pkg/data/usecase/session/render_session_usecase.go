package session

import (
	"net/http"

	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
)

type renderSessionUseCase struct{}

func (c *renderSessionUseCase) Name() string {
	return "render_session_usecase"
}

func (c *renderSessionUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpWriter, errWriter := ctxAware.GetWriter()
	session, errSession := ctxAware.GetSession(TagSession)

	if errWriter != nil || errSession != nil {
		return httpwrapper.CreateInternalError()
	}

	httpWriter.WriteJson(http.StatusOK, session)
	return nil
}

func CreateRenderSessionUseCase() pipeline.Step {
	return &renderSessionUseCase{}
}

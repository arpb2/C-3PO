package user

import (
	"net/http"

	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"

	"github.com/saantiaguilera/go-pipeline"
)

type renderEmptyUseCase struct{}

func (c *renderEmptyUseCase) Name() string {
	return "render_empty_usecase"
}

func (c *renderEmptyUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpWriter, err := ctxAware.GetWriter()

	if err != nil {
		return err
	}

	httpWriter.WriteStatus(http.StatusOK)
	return nil
}

func CreateRenderEmptyUseCase() pipeline.Step {
	return &renderEmptyUseCase{}
}

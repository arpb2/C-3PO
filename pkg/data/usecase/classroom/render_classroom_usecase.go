package classroom

import (
	"net/http"

	ctxaware "github.com/arpb2/C-3PO/pkg/domain/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
)

type renderClassroomUseCase struct{}

func (c *renderClassroomUseCase) Name() string {
	return "render_classroom_usecase"
}

func (c *renderClassroomUseCase) Run(ctx pipeline.Context) error {
	ctxAware := ctxaware.CreateContextAware(ctx)

	httpWriter, errWriter := ctxAware.GetWriter()
	cr, errCr := ctxAware.GetClassroom(TagClassroom)

	if errWriter != nil || errCr != nil {
		return httpwrapper.CreateInternalError()
	}

	httpWriter.WriteJson(http.StatusOK, cr)
	return nil
}

func CreateRenderClassroomUseCase() pipeline.Step {
	return &renderClassroomUseCase{}
}

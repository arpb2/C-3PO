package classroom

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model/classroom"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type fetchClassroomUseCase struct{}

func (c *fetchClassroomUseCase) Name() string {
	return "fetch_classroom_usecase"
}

func (c *fetchClassroomUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	var cr classroom.Classroom
	if err := httpReader.ReadBody(&cr); err != nil {
		return http.CreateBadRequestError("malformed body")
	}

	ctx.Set(TagClassroom, cr)
	return nil
}

func CreateFetchClassroomUseCase() pipeline.Step {
	return &fetchClassroomUseCase{}
}

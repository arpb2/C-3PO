package classroom

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/classroom"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type updateClassroomUseCase struct {
	repository classroom.Repository
}

func (c *updateClassroomUseCase) Name() string {
	return "update_classroom_usecase"
}

func (c *updateClassroomUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	cr, errCr := ctxAware.GetClassroom(TagClassroom)
	crID, existsID := ctx.GetUInt(TagClassroomID)

	if errCr != nil || !existsID {
		return http.CreateInternalError()
	}

	cr.Id = crID

	class, err := c.repository.UpdateClassroom(cr)

	if err != nil {
		return err
	}

	ctx.Set(TagClassroom, class)
	return nil
}

func CreateUpdateClassroomUseCase(repository classroom.Repository) pipeline.Step {
	return &updateClassroomUseCase{
		repository: repository,
	}
}

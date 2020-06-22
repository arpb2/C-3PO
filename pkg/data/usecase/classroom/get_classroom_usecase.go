package classroom

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/classroom"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/saantiaguilera/go-pipeline"
)

type getClassroomUseCase struct {
	repository classroom.Repository
}

func (c *getClassroomUseCase) Name() string {
	return "get_classroom_usecase"
}

func (c *getClassroomUseCase) Run(ctx pipeline.Context) error {
	id, exists := ctx.GetUInt(TagClassroomID)

	if !exists {
		return http.CreateInternalError()
	}

	cr, err := c.repository.GetClassroom(id)

	if err != nil {
		return err
	}

	ctx.Set(TagClassroom, cr)
	return nil
}

func CreateGetClassroomUseCase(repository classroom.Repository) pipeline.Step {
	return &getClassroomUseCase{
		repository: repository,
	}
}

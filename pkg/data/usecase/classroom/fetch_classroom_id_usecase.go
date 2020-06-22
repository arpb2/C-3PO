package classroom

import (
	"fmt"
	"strconv"

	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/pkg/domain/http"
)

type fetchClassroomIdUseCase struct {
	ClassroomIdParam string
}

func (c *fetchClassroomIdUseCase) Name() string {
	return fmt.Sprintf("fetch_%s_usecase", c.ClassroomIdParam)
}

func (c *fetchClassroomIdUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	id := httpReader.GetParameter(c.ClassroomIdParam)

	if id == "" {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' empty", c.ClassroomIdParam))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' malformed, expecting a positive number", c.ClassroomIdParam))
	}

	ctx.Set(TagClassroomID, uint(idUint))
	return nil
}

func CreateFetchClassroomIdUseCase(classroomIDParam string) pipeline.Step {
	return &fetchClassroomIdUseCase{
		ClassroomIdParam: classroomIDParam,
	}
}

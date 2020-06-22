package classroom

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/classroom"
	classroom2 "github.com/arpb2/C-3PO/pkg/data/usecase/classroom"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateGetHandler(paramClassroomID string, exec pipeline.HttpPipeline, classroomRepository classroom.Repository) http.Handler {
	fetchClassroomIdUseCase := classroom2.CreateFetchClassroomIdUseCase(paramClassroomID)
	repositoryUseCase := classroom2.CreateGetClassroomUseCase(classroomRepository)
	renderUseCase := classroom2.CreateRenderClassroomUseCase()

	graph := gopipeline.CreateSequentialStage(
		fetchClassroomIdUseCase,
		repositoryUseCase,
		renderUseCase,
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}

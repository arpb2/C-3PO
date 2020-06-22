package classroom

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/classroom"
	classroom2 "github.com/arpb2/C-3PO/pkg/data/usecase/classroom"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePutHandler(paramClassroomID string, exec pipeline.HttpPipeline, classroomRepository classroom.Repository) http.Handler {
	fetchClassroomIDUseCase := classroom2.CreateFetchClassroomIdUseCase(paramClassroomID)
	fetchClassroomUseCase := classroom2.CreateFetchClassroomUseCase()
	repositoryUseCase := classroom2.CreateUpdateClassroomUseCase(classroomRepository)
	renderUseCase := classroom2.CreateRenderClassroomUseCase()

	graph := gopipeline.CreateSequentialGroup(
		gopipeline.CreateConcurrentStage(
			fetchClassroomIDUseCase,
			fetchClassroomUseCase,
		),
		gopipeline.CreateSequentialStage(
			repositoryUseCase,
			renderUseCase,
		),
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}

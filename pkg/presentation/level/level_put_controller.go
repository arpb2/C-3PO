package level

import (
	level2 "github.com/arpb2/C-3PO/pkg/data/repository/level"
	"github.com/arpb2/C-3PO/pkg/data/usecase/level"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePutHandler(paramLevelId string, exec pipeline.HttpPipeline, levelRepository level2.Repository) http.Handler {
	fetchLevelIdUseCase := level.CreateFetchLevelIdUseCase(paramLevelId)
	fetchLevelUseCase := level.CreateFetchLevelUseCase()
	repositoryUseCase := level.CreateWriteLevelUseCase(levelRepository)
	renderUseCase := level.CreateRenderLevelUseCase()

	graph := gopipeline.CreateSequentialGroup(
		gopipeline.CreateConcurrentStage(
			fetchLevelIdUseCase,
			fetchLevelUseCase,
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

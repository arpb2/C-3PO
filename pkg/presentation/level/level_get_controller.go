package level

import (
	level2 "github.com/arpb2/C-3PO/pkg/data/repository/level"
	"github.com/arpb2/C-3PO/pkg/data/usecase/level"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateGetHandler(paramLevelId string, exec pipeline.HttpPipeline, levelRepository level2.Repository) http.Handler {
	fetchLevelIdUseCase := level.CreateFetchLevelIdUseCase(paramLevelId)
	repositoryUseCase := level.CreateGetLevelUseCase(levelRepository)
	renderUseCase := level.CreateRenderLevelUseCase()

	graph := gopipeline.CreateSequentialStage(
		fetchLevelIdUseCase,
		repositoryUseCase,
		renderUseCase,
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}

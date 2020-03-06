package user

import (
	user2 "github.com/arpb2/C-3PO/pkg/data/repository/user"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateGetLevelHandler(paramUserId, paramLevelId string, exec pipeline.HttpPipeline, userLevelRepository user2.LevelRepository) http.Handler {
	fetchUserIdUseCase := user.CreateFetchUserIdUseCase(paramUserId)
	fetchLevelIdUseCase := user.CreateFetchLevelIdUseCase(paramLevelId)
	repositoryUseCase := user.CreateGetUserLevelUseCase(userLevelRepository)
	renderUseCase := user.CreateRenderUserLevelUseCase()

	graph := gopipeline.CreateSequentialGroup(
		gopipeline.CreateConcurrentStage(
			fetchUserIdUseCase,
			fetchLevelIdUseCase,
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

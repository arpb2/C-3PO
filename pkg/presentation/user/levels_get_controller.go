package user

import (
	user2 "github.com/arpb2/C-3PO/pkg/data/repository/user"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateGetLevelsHandler(paramUserId string, exec pipeline.HttpPipeline, userLevelRepository user2.LevelRepository) http.Handler {
	fetchUserIdUseCase := user.CreateFetchUserIdUseCase(paramUserId)
	repositoryUseCase := user.CreateGetUserLevelsUseCase(userLevelRepository)
	renderUseCase := user.CreateRenderUserLevelsUseCase()

	graph := gopipeline.CreateSequentialStage(
		fetchUserIdUseCase,
		repositoryUseCase,
		renderUseCase,
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}

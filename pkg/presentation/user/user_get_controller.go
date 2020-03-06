package user

import (
	user2 "github.com/arpb2/C-3PO/pkg/data/repository/user"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateGetUserHandler(paramUserId string, exec pipeline.HttpPipeline, userRepository user2.Repository) http.Handler {
	fetchUserIdUseCase := user.CreateFetchUserIdUseCase(paramUserId)
	repositoryUseCase := user.CreateGetUserUseCase(userRepository)
	renderUseCase := user.CreateRenderUserUseCase()

	graph := gopipeline.CreateSequentialStage(
		fetchUserIdUseCase,
		repositoryUseCase,
		renderUseCase,
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}

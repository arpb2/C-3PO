package user

import (
	user2 "github.com/arpb2/C-3PO/pkg/data/repository/user"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user/validation"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePutUserHandler(
	paramUserId string,
	exec pipeline.HttpPipeline,
	userRepository user2.Repository,
	validations []validation.Validation,
) http.Handler {
	fetchUserIdUseCase := user.CreateFetchUserIdUseCase(paramUserId)
	fetchUserUseCase := user.CreateFetchAuthenticatedUserUseCase()
	validateUseCase := user.CreateValidateUserParametersUseCase(validations)
	repositoryUseCase := user.CreateUpdateUserUseCase(userRepository)
	renderUseCase := user.CreateRenderUserUseCase()

	graph := gopipeline.CreateSequentialGroup(
		gopipeline.CreateConcurrentGroup(
			gopipeline.CreateSequentialStage(
				fetchUserIdUseCase,
			),
			gopipeline.CreateSequentialStage(
				fetchUserUseCase,
				validateUseCase,
			),
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

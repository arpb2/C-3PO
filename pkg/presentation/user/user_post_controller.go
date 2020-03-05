package user

import (
	user2 "github.com/arpb2/C-3PO/pkg/data/repository/user"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user/validation"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePostUserHandler(
	exec pipeline.HttpPipeline,
	userRepository user2.Repository,
	validations []validation.Validation,
) http.Handler {
	fetchUserUseCase := user.CreateFetchAuthenticatedUserUseCase()
	validateUseCase := user.CreateValidateUserParametersUseCase(validations)
	createUserUseCase := user.CreateCreateUserUseCase(userRepository)
	renderUseCase := user.CreateRenderUserUseCase()

	graph := gopipeline.CreateSequentialStage(
		fetchUserUseCase,
		validateUseCase,
		createUserUseCase,
		renderUseCase,
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}

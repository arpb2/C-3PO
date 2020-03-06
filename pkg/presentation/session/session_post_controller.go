package session

import (
	session2 "github.com/arpb2/C-3PO/pkg/data/repository/session"
	"github.com/arpb2/C-3PO/pkg/data/usecase/session"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user/validation"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePostHandler(
	executor pipeline.HttpPipeline,
	tokenHandler session2.TokenRepository,
	repository session2.CredentialRepository,
	validations []validation.Validation,
) http.Handler {
	fetchUserUseCase := user.CreateFetchAuthenticatedUserUseCase()
	validateParamsUseCase := user.CreateValidateUserParametersUseCase(validations)
	authenticateUseCase := session.CreateAuthenticateUseCase(repository)
	createSessionUseCase := session.CreateCreateSessionUseCase(tokenHandler)
	renderUseCase := session.CreateRenderSessionUseCase()

	graph := gopipeline.CreateSequentialStage(
		fetchUserUseCase,
		validateParamsUseCase,
		authenticateUseCase,
		createSessionUseCase,
		renderUseCase,
	)

	return func(ctx *http.Context) {
		executor.Run(ctx, graph)
	}
}

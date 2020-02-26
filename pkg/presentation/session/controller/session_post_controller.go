package controller

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	credentialrepository "github.com/arpb2/C-3PO/pkg/domain/session/repository"
	"github.com/arpb2/C-3PO/pkg/presentation/session/command"
	usercommand "github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePostController(executor pipeline.HttpPipeline,
	tokenHandler credentialrepository.TokenRepository,
	repository credentialrepository.CredentialRepository,
	validations []validation.Validation) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/session",
		Body:   CreatePostBody(executor, tokenHandler, repository, validations),
	}
}

func CreatePostBody(executor pipeline.HttpPipeline, tokenHandler credentialrepository.TokenRepository, repository credentialrepository.CredentialRepository, validations []validation.Validation) http.Handler {
	fetchUserCommand := usercommand.CreateFetchAuthenticatedUserCommand()
	validateParamsCommand := usercommand.CreateValidateUserParametersCommand(validations)
	authenticateCommand := command.CreateAuthenticateCommand(repository)
	createSessionCommand := command.CreateCreateSessionCommand(tokenHandler)
	renderCommand := command.CreateRenderSessionCommand()

	graph := gopipeline.CreateSequentialStage(
		fetchUserCommand,
		validateParamsCommand,
		authenticateCommand,
		createSessionCommand,
		renderCommand,
	)

	return func(ctx *http.Context) {
		executor.Run(ctx, graph)
	}
}

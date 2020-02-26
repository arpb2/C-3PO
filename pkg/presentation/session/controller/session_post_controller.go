package controller

import (
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/controller"
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"
	credentialservice "github.com/arpb2/C-3PO/pkg/domain/session/service"
	"github.com/arpb2/C-3PO/pkg/domain/session/token"
	"github.com/arpb2/C-3PO/pkg/presentation/session/command"
	usercommand "github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePostController(executor pipeline.HttpPipeline,
	tokenHandler token.Handler,
	service credentialservice.Service,
	validations []validation.Validation) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/session",
		Body:   CreatePostBody(executor, tokenHandler, service, validations),
	}
}

func CreatePostBody(executor pipeline.HttpPipeline, tokenHandler token.Handler, service credentialservice.Service, validations []validation.Validation) http.Handler {
	fetchUserCommand := usercommand.CreateFetchAuthenticatedUserCommand()
	validateParamsCommand := usercommand.CreateValidateUserParametersCommand(validations)
	authenticateCommand := command.CreateAuthenticateCommand(service)
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

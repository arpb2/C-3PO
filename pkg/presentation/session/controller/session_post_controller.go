package controller

import (
	"github.com/arpb2/C-3PO/pkg/domain/auth"
	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	credentialservice "github.com/arpb2/C-3PO/pkg/domain/service/credential"
	"github.com/arpb2/C-3PO/pkg/presentation/session/command"
	command2 "github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePostController(executor pipeline.HttpPipeline,
	tokenHandler auth.TokenHandler,
	service credentialservice.Service,
	validations []validation.Validation) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/session",
		Body:   CreatePostBody(executor, tokenHandler, service, validations),
	}
}

func CreatePostBody(executor pipeline.HttpPipeline, tokenHandler auth.TokenHandler, service credentialservice.Service, validations []validation.Validation) http.Handler {
	fetchUserCommand := command2.CreateFetchAuthenticatedUserCommand()
	validateParamsCommand := command2.CreateValidateUserParametersCommand(validations)
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

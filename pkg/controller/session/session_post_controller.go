package session

import (
	"github.com/arpb2/C-3PO/api/auth"
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	credentialservice "github.com/arpb2/C-3PO/api/service/credential"
	sessioncommand "github.com/arpb2/C-3PO/pkg/command/session"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	uservalidation "github.com/arpb2/C-3PO/pkg/validation/user"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
)

func CreatePostController(executor pipeline.HttpPipeline,
	tokenHandler auth.TokenHandler,
	service credentialservice.Service,
	validations []uservalidation.Validation) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/session",
		Body:   CreatePostBody(executor, tokenHandler, service, validations),
	}
}

func CreatePostBody(executor pipeline.HttpPipeline, tokenHandler auth.TokenHandler, service credentialservice.Service, validations []uservalidation.Validation) http.Handler {
	return func(ctx *http.Context) {
		fetchUserCommand := usercommand.CreateFetchAuthenticatedUserCommand(ctx)
		validateParamsCommand := usercommand.CreateValidateParametersCommand(ctx, fetchUserCommand.OutputStream, validations)
		authenticateCommand := sessioncommand.CreateAuthenticateCommand(ctx, service, validateParamsCommand.OutputStream)
		createTokenCommand := sessioncommand.CreateCreateTokenCommand(ctx, tokenHandler, authenticateCommand.OutputStream)
		renderCommand := sessioncommand.CreateRenderSessionCommand(ctx, createTokenCommand.OutputStream)

		graph := sequential.CreateSequentialStage(
			fetchUserCommand,
			validateParamsCommand,
			authenticateCommand,
			createTokenCommand,
			renderCommand,
		)

		executor.Run(ctx, graph)
	}
}

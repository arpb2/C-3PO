package session_controller

import (
	"github.com/arpb2/C-3PO/api/auth"
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/pipeline"
	credential_service "github.com/arpb2/C-3PO/api/service/credential"
	session_command "github.com/arpb2/C-3PO/pkg/command/session"
	user_command "github.com/arpb2/C-3PO/pkg/command/user"
	user_validation "github.com/arpb2/C-3PO/pkg/validation/user"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
)

func CreatePostController(executor pipeline.HttpPipeline,
	tokenHandler auth.TokenHandler,
	service credential_service.Service,
	validations []user_validation.Validation) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/session",
		Body:   CreatePostBody(executor, tokenHandler, service, validations),
	}
}

func CreatePostBody(executor pipeline.HttpPipeline, tokenHandler auth.TokenHandler, service credential_service.Service, validations []user_validation.Validation) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		fetchUserCommand := user_command.CreateFetchAuthenticatedUserCommand(ctx)
		validateParamsCommand := user_command.CreateValidateParametersCommand(ctx, fetchUserCommand.OutputStream, validations)
		authenticateCommand := session_command.CreateAuthenticateCommand(ctx, service, validateParamsCommand.OutputStream)
		createTokenCommand := session_command.CreateCreateTokenCommand(ctx, tokenHandler, authenticateCommand.OutputStream)
		renderCommand := session_command.CreateRenderSessionCommand(ctx, createTokenCommand.OutputStream)

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

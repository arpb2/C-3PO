package session

import (
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/command/session_command"
	"github.com/arpb2/C-3PO/src/api/command/user_command"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_validation"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
)

func CreatePostController(executor executor.HttpPipeline,
						  tokenHandler auth.TokenHandler,
	                      service service.CredentialService,
	                      validations []session_validation.Validation) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/session",
		Body:   PostBody{
			Executor:     executor,
			TokenHandler: tokenHandler,
			Service:      service,
			Validations:  validations,
		}.Method,
	}
}

type PostBody struct {
	Executor     executor.HttpPipeline
	TokenHandler auth.TokenHandler
	Service      service.CredentialService

	Validations    []session_validation.Validation
}

func (b PostBody) Method(ctx *http_wrapper.Context) {
	fetchUserCommand := user_command.CreateFetchAuthenticatedUserCommand(ctx)
	validateParamsCommand := session_command.CreateValidateParametersCommand(ctx, fetchUserCommand.OutputStream, b.Validations)
	authenticateCommand := session_command.CreateAuthenticateCommand(ctx, b.Service, validateParamsCommand.OutputStream)
	createTokenCommand := session_command.CreateCreateTokenCommand(ctx, b.TokenHandler, authenticateCommand.OutputStream)
	renderCommand := session_command.CreateRenderSessionCommand(ctx, createTokenCommand.OutputStream)

	graph := sequential.CreateSequentialStage(
		fetchUserCommand,
		validateParamsCommand,
		authenticateCommand,
		createTokenCommand,
		renderCommand,
	)

	b.Executor.Run(ctx, graph)
}

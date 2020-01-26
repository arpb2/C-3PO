package session_controller

import (
	"github.com/arpb2/C-3PO/api/auth"
	controller2 "github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/pipeline"
	"github.com/arpb2/C-3PO/api/service/credential"
	session_command "github.com/arpb2/C-3PO/internal/command/session"
	user_command "github.com/arpb2/C-3PO/internal/command/user"
	user_validation "github.com/arpb2/C-3PO/internal/validation/user"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
)

func CreatePostController(executor pipeline.HttpPipeline,
	tokenHandler auth.TokenHandler,
	service credential_service.Service,
	validations []user_validation.Validation) controller2.Controller {
	return controller2.Controller{
		Method: "POST",
		Path:   "/session",
		Body: PostBody{
			Executor:     executor,
			TokenHandler: tokenHandler,
			Service:      service,
			Validations:  validations,
		}.Method,
	}
}

type PostBody struct {
	Executor     pipeline.HttpPipeline
	TokenHandler auth.TokenHandler
	Service      credential_service.Service

	Validations []user_validation.Validation
}

func (b PostBody) Method(ctx *http_wrapper.Context) {
	fetchUserCommand := user_command.CreateFetchAuthenticatedUserCommand(ctx)
	validateParamsCommand := user_command.CreateValidateParametersCommand(ctx, fetchUserCommand.OutputStream, b.Validations)
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

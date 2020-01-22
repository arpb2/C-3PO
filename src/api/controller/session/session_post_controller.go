package session

import (
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_command"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_validation"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_command"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
	"net/http"
)

func CreatePostController(executor executor.Executor,
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
	Executor     executor.Executor
	TokenHandler auth.TokenHandler
	Service      service.CredentialService

	Validations    []session_validation.Validation
}

func (b PostBody) Method(ctx *http_wrapper.Context) {
	fetchUserCommand := user_command.CreateFetchAuthenticatedUserCommand(ctx)
	validateParamsCommand := session_command.CreateValidateParametersCommand(ctx, fetchUserCommand.OutputStream, b.Validations)
	authenticateCommand := session_command.CreateAuthenticateCommand(ctx, b.Service, validateParamsCommand.OutputStream)
	createTokenCommand := session_command.CreateCreateTokenCommand(ctx, b.TokenHandler, authenticateCommand.OutputStream)

	commands := []executor.Command{
		fetchUserCommand,
		validateParamsCommand,
		authenticateCommand,
		createTokenCommand,
	}

	if err := controller.BatchRun(b.Executor, commands, ctx); err == nil {
		ctx.WriteJson(http.StatusOK, http_wrapper.Json{
			"token": <-createTokenCommand.OutputStream,
		})
	}
}

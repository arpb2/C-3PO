package session

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_command"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_validation"
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
	fetchUserCommand := session_command.CreateFetchUserCommand(ctx)
	validateParamsCommand := session_command.CreateValidateParametersCommand(ctx, fetchUserCommand.OutputStream, b.Validations)
	authenticateCommand := session_command.CreateAuthenticateCommand(ctx, b.Service, validateParamsCommand.Stream)
	createTokenCommand := session_command.CreateTokenCommand(ctx, b.TokenHandler, authenticateCommand.Stream)

	commands := []executor.Command{
		fetchUserCommand,
		validateParamsCommand,
		authenticateCommand,
		createTokenCommand,
	}

	for _, command := range commands {
		err := b.Executor.Do(command)

		if ctx.IsAborted() {
			return
		}

		if err != nil {
			fmt.Print(err.Error())
			controller.Halt(ctx, http.StatusInternalServerError, "internal error")
			return
		}
	}

	token := <-createTokenCommand.OutputStream
	ctx.WriteJson(http.StatusOK, http_wrapper.Json{
		"token": token,
	})
}

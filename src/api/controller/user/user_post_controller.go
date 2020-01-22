package user

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_command"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_validation"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
	"net/http"
)

func CreatePostController(exec executor.Executor, validations []user_validation.Validation, userService service.UserService) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/users",
		Body: CreatePostBody(exec, validations, userService),
	}
}

func CreatePostBody(exec executor.Executor, validations []user_validation.Validation, userService service.UserService) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		fetchUserCommand := user_command.CreateFetchAuthenticatedUserCommand(ctx)
		validateCommand := user_command.CreateValidateParametersCommand(ctx, fetchUserCommand.OutputStream, validations)
		serviceCommand := user_command.CreateCreateUserCommand(ctx, userService, fetchUserCommand.OutputStream)

		commands := []executor.Command{
			fetchUserCommand,
			validateCommand,
			serviceCommand,
		}

		if err := controller.BatchRun(exec, commands, ctx); err == nil {
			ctx.WriteJson(http.StatusOK, <-serviceCommand.OutputStream)
		}
	}
}
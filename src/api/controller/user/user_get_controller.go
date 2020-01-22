package user

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_command"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
	"net/http"
)

func CreateGetController(executor executor.Executor,
						 authMiddleware http_wrapper.Handler,
						 userService service.UserService) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   "/users/:user_id",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body:   CreateGetBody(executor, userService),
	}
}

func CreateGetBody(exec executor.Executor, userService service.UserService) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		fetchUserIdCommand := user_command.CreateFetchUserIdCommand(ctx)
		serviceCommand := user_command.CreateGetUserCommand(ctx, userService, fetchUserIdCommand.OutputStream)

		commands := []executor.Command{
			fetchUserIdCommand,
			serviceCommand,
		}

		for _, command := range commands {
			err := exec.Do(command)

			if ctx.IsAborted() {
				return
			}

			if err != nil {
				fmt.Print(err.Error())
				controller.Halt(ctx, http.StatusInternalServerError, "internal error")
				return
			}
		}

		user := <-serviceCommand.OutputStream
		ctx.WriteJson(http.StatusOK, user)
	}
}
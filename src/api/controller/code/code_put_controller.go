package code

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/code/code_command"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_command"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
	"net/http"
)

func CreatePutController(exec executor.Executor, authMiddleware http_wrapper.Handler, codeService service.CodeService) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   "/users/:user_id/codes/:code_id",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body:   CreatePutBody(exec, codeService),
	}
}

func CreatePutBody(exec executor.Executor, codeService service.CodeService) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		fetchUserIdCommand := user_command.CreateFetchUserIdCommand(ctx)
		fetchCodeCommand := code_command.CreateFetchCodeCommand(ctx)
		fetchCodeIdCommand := code_command.CreateFetchCodeIdCommand(ctx)

		codeChan := make(chan *model.Code, 1)
		go func() {
			defer close(codeChan)
			userId, openUserIdChan := <-fetchUserIdCommand.OutputStream
			codeId, openCodeIdChan := <-fetchCodeIdCommand.OutputStream
			code, openCodeChan := <-fetchCodeCommand.OutputStream

			if !openUserIdChan && !openCodeIdChan && !openCodeChan {
				return
			}

			codeChan <- &model.Code{
				Id:     codeId,
				UserId: userId,
				Code:   code,
			}
		}()

		serviceCommand := code_command.CreateReplaceCodeCommand(ctx, codeService, codeChan)

		commands := []executor.Command{
			fetchUserIdCommand,
			fetchCodeCommand,
			serviceCommand,
		}

		if err := controller.BatchRun(exec, commands, ctx); err == nil {
			code := <-serviceCommand.OutputStream
			ctx.WriteJson(http.StatusOK, *code)
		}
	}
}

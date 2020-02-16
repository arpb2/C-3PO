package code

import (
	"fmt"

	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	codeservice "github.com/arpb2/C-3PO/api/service/code"
	codecommand "github.com/arpb2/C-3PO/pkg/command/code"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePutController(exec pipeline.HttpPipeline, authMiddleware http.Handler, codeService codeservice.Service) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   fmt.Sprintf("/users/:%s/codes/:%s", ParamUserId, ParamCodeId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreatePutBody(exec, codeService),
	}
}

func CreatePutBody(exec pipeline.HttpPipeline, codeService codeservice.Service) http.Handler {
	fetchUserIdCommand := usercommand.CreateFetchUserIdCommand()
	fetchCodeCommand := codecommand.CreateFetchCodeCommand()
	fetchCodeIdCommand := codecommand.CreateFetchCodeIdCommand()
	serviceCommand := codecommand.CreateReplaceCodeCommand(codeService)
	renderCommand := codecommand.CreateRenderCodeCommand()

	graph := gopipeline.CreateSequentialGroup(
		gopipeline.CreateConcurrentStage(
			fetchUserIdCommand,
			fetchCodeCommand,
			fetchCodeIdCommand,
		),
		gopipeline.CreateSequentialStage(
			serviceCommand,
			renderCommand,
		),
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}

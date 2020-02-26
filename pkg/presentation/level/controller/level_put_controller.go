package controller

import (
	"fmt"

	"github.com/arpb2/C-3PO/pkg/presentation/level"

	"github.com/arpb2/C-3PO/pkg/domain/level/service"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	"github.com/arpb2/C-3PO/pkg/presentation/level/command"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePutController(exec pipeline.HttpPipeline, authMiddleware http.Handler, levelService service.Service) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   fmt.Sprintf("/levels/:%s", level.ParamLevelId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreatePutBody(exec, levelService),
	}
}

func CreatePutBody(exec pipeline.HttpPipeline, levelService service.Service) http.Handler {
	fetchLevelIdCommand := command.CreateFetchLevelIdCommand()
	fetchLevelCommand := command.CreateFetchLevelCommand()
	serviceCommand := command.CreateWriteLevelCommand(levelService)
	renderCommand := command.CreateRenderLevelCommand()

	graph := gopipeline.CreateSequentialGroup(
		gopipeline.CreateConcurrentStage(
			fetchLevelIdCommand,
			fetchLevelCommand,
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

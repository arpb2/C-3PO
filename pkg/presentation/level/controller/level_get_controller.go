package controller

import (
	"fmt"

	controller3 "github.com/arpb2/C-3PO/pkg/domain/level/controller"
	"github.com/arpb2/C-3PO/pkg/domain/level/service"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	"github.com/arpb2/C-3PO/pkg/presentation/level/command"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateGetController(exec pipeline.HttpPipeline, levelService service.Service) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   fmt.Sprintf("/levels/:%s", controller3.ParamLevelId),
		Body:   CreateGetBody(exec, levelService),
	}
}

func CreateGetBody(exec pipeline.HttpPipeline, levelService service.Service) http.Handler {
	fetchLevelIdCommand := command.CreateFetchLevelIdCommand()
	serviceCommand := command.CreateGetLevelCommand(levelService)
	renderCommand := command.CreateRenderLevelCommand()

	graph := gopipeline.CreateSequentialStage(
		fetchLevelIdCommand,
		serviceCommand,
		renderCommand,
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}

package controller

import (
	"fmt"

	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	levelservice "github.com/arpb2/C-3PO/pkg/domain/service/level"
	"github.com/arpb2/C-3PO/pkg/presentation/level/command"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateGetController(exec pipeline.HttpPipeline, levelService levelservice.Service) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   fmt.Sprintf("/levels/:%s", controller.ParamLevelId),
		Body:   CreateGetBody(exec, levelService),
	}
}

func CreateGetBody(exec pipeline.HttpPipeline, levelService levelservice.Service) http.Handler {
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

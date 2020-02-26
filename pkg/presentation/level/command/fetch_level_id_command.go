package command

import (
	"fmt"
	"strconv"

	controller3 "github.com/arpb2/C-3PO/pkg/domain/level/controller"

	httppipeline "github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
)

type fetchLevelIdCommand struct{}

func (c *fetchLevelIdCommand) Name() string {
	return "fetch_level_id_command"
}

func (c *fetchLevelIdCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	levelId := httpReader.GetParameter(controller3.ParamLevelId)

	if levelId == "" {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' empty", controller3.ParamLevelId))
	}

	levelIdUint, err := strconv.ParseUint(levelId, 10, 64)

	if err != nil {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' malformed, expecting a positive number", controller3.ParamLevelId))
	}

	ctx.Set(TagLevelId, uint(levelIdUint))
	return nil
}

func CreateFetchLevelIdCommand() pipeline.Step {
	return &fetchLevelIdCommand{}
}

package command

import (
	"fmt"
	"strconv"

	controller2 "github.com/arpb2/C-3PO/pkg/domain/controller"

	httppipeline "github.com/arpb2/C-3PO/pkg/infra/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/pkg/domain/http"
)

type fetchLevelIdCommand struct {
	context *http.Context
}

func (c *fetchLevelIdCommand) Name() string {
	return fmt.Sprintf("fetch_%s_command", controller2.ParamLevelId)
}

func (c *fetchLevelIdCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	levelId := httpReader.GetParameter(controller2.ParamLevelId)

	if levelId == "" {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' empty", controller2.ParamLevelId))
	}

	levelIdUint, err := strconv.ParseUint(levelId, 10, 64)

	if err != nil {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' malformed, expecting a positive number", controller2.ParamLevelId))
	}

	ctx.Set(TagLevelId, uint(levelIdUint))
	return nil
}

func CreateFetchLevelIdCommand() pipeline.Step {
	return &fetchLevelIdCommand{}
}

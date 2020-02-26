package command

import (
	"fmt"
	"strconv"

	"github.com/arpb2/C-3PO/pkg/presentation/level"

	"github.com/arpb2/C-3PO/pkg/presentation/level/command"

	httppipeline "github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
)

type fetchLevelIdCommand struct {
	context *http.Context
}

func (c *fetchLevelIdCommand) Name() string {
	return fmt.Sprintf("fetch_%s_command", level.ParamLevelId)
}

func (c *fetchLevelIdCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	levelId := httpReader.GetParameter(level.ParamLevelId)

	if levelId == "" {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' empty", level.ParamLevelId))
	}

	levelIdUint, err := strconv.ParseUint(levelId, 10, 64)

	if err != nil {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' malformed, expecting a positive number", level.ParamLevelId))
	}

	ctx.Set(command.TagLevelId, uint(levelIdUint))
	return nil
}

func CreateFetchLevelIdCommand() pipeline.Step {
	return &fetchLevelIdCommand{}
}

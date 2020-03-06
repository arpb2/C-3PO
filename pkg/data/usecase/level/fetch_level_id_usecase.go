package level

import (
	"fmt"
	"strconv"

	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/pkg/domain/http"
)

type fetchLevelIdUseCase struct {
	ParamLevelId string
}

func (c *fetchLevelIdUseCase) Name() string {
	return fmt.Sprintf("fetch_%s_usecase", c.ParamLevelId)
}

func (c *fetchLevelIdUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	levelId := httpReader.GetParameter(c.ParamLevelId)

	if levelId == "" {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' empty", c.ParamLevelId))
	}

	levelIdUint, err := strconv.ParseUint(levelId, 10, 64)

	if err != nil {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' malformed, expecting a positive number", c.ParamLevelId))
	}

	ctx.Set(TagLevelId, uint(levelIdUint))
	return nil
}

func CreateFetchLevelIdUseCase(paramLevelId string) pipeline.Step {
	return &fetchLevelIdUseCase{
		ParamLevelId: paramLevelId,
	}
}

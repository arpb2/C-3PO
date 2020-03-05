package user

import (
	"fmt"
	"strconv"

	level2 "github.com/arpb2/C-3PO/pkg/data/usecase/level"

	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/pkg/domain/http"
)

type fetchLevelIdUseCase struct {
	context      *http.Context
	LevelIdParam string
}

func (c *fetchLevelIdUseCase) Name() string {
	return fmt.Sprintf("fetch_%s_usecase", c.LevelIdParam)
}

func (c *fetchLevelIdUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	levelId := httpReader.GetParameter(c.LevelIdParam)

	if levelId == "" {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' empty", c.LevelIdParam))
	}

	levelIdUint, err := strconv.ParseUint(levelId, 10, 64)

	if err != nil {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' malformed, expecting a positive number", c.LevelIdParam))
	}

	ctx.Set(level2.TagLevelId, uint(levelIdUint))
	return nil
}

func CreateFetchLevelIdUseCase(levelIdParam string) pipeline.Step {
	return &fetchLevelIdUseCase{
		LevelIdParam: levelIdParam,
	}
}

package level

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/level"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type writeLevelUseCase struct {
	repository level.Repository
}

func (c *writeLevelUseCase) Name() string {
	return "write_level_usecase"
}

func (c *writeLevelUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	levelId, existsLevelId := ctx.GetUInt(TagLevelId)
	levelData, existsData := ctxAware.GetLevel(TagLevel)

	if !existsLevelId || existsData != nil {
		return http.CreateInternalError()
	}

	levelData.Id = levelId
	lvl, err := c.repository.StoreLevel(levelData)

	if err != nil {
		return err
	}

	ctx.Set(TagLevel, lvl)
	return nil
}

func CreateWriteLevelUseCase(repository level.Repository) pipeline.Step {
	return &writeLevelUseCase{
		repository: repository,
	}
}

package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	repository2 "github.com/arpb2/C-3PO/pkg/domain/level/repository"
	"github.com/saantiaguilera/go-pipeline"
)

type writeLevelCommand struct {
	repository repository2.LevelRepository
}

func (c *writeLevelCommand) Name() string {
	return "write_level_command"
}

func (c *writeLevelCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	levelId, existsLevelId := ctx.GetUInt(TagLevelId)
	levelData, existsData := ctxAware.GetLevel(TagLevel)

	if !existsLevelId || existsData != nil {
		return http.CreateInternalError()
	}

	levelData.Id = levelId
	level, err := c.repository.StoreLevel(levelData)

	if err != nil {
		return err
	}

	ctx.Set(TagLevel, level)
	return nil
}

func CreateWriteLevelCommand(repository repository2.LevelRepository) pipeline.Step {
	return &writeLevelCommand{
		repository: repository,
	}
}

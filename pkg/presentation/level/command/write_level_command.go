package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"
	service2 "github.com/arpb2/C-3PO/pkg/domain/level/service"
	"github.com/saantiaguilera/go-pipeline"
)

type writeLevelCommand struct {
	service service2.Service
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
	level, err := c.service.StoreLevel(levelData)

	if err != nil {
		return err
	}

	ctx.Set(TagLevel, level)
	return nil
}

func CreateWriteLevelCommand(service service2.Service) pipeline.Step {
	return &writeLevelCommand{
		service: service,
	}
}

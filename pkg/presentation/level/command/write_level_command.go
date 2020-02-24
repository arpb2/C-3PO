package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
	levelservice "github.com/arpb2/C-3PO/pkg/domain/service/level"
	pipeline2 "github.com/arpb2/C-3PO/pkg/infra/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type writeLevelCommand struct {
	service levelservice.Service
}

func (c *writeLevelCommand) Name() string {
	return "write_level_command"
}

func (c *writeLevelCommand) Run(ctx pipeline.Context) error {
	ctxAware := pipeline2.CreateContextAware(ctx)

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

func CreateWriteLevelCommand(service levelservice.Service) pipeline.Step {
	return &writeLevelCommand{
		service: service,
	}
}

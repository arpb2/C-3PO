package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	service2 "github.com/arpb2/C-3PO/pkg/domain/level/service"
	"github.com/saantiaguilera/go-pipeline"
)

type getLevelCommand struct {
	service service2.Service
}

func (c *getLevelCommand) Name() string {
	return "get_level_command"
}

func (c *getLevelCommand) Run(ctx pipeline.Context) error {
	levelId, existsUserId := ctx.GetUInt(TagLevelId)

	if !existsUserId {
		return http.CreateInternalError()
	}

	level, err := c.service.GetLevel(levelId)

	if err != nil {
		return err
	}

	ctx.Set(TagLevel, level)
	return nil
}

func CreateGetLevelCommand(service service2.Service) pipeline.Step {
	return &getLevelCommand{
		service: service,
	}
}

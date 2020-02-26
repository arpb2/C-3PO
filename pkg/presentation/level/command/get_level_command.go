package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	repository2 "github.com/arpb2/C-3PO/pkg/domain/level/repository"
	"github.com/saantiaguilera/go-pipeline"
)

type getLevelCommand struct {
	repository repository2.LevelRepository
}

func (c *getLevelCommand) Name() string {
	return "get_level_command"
}

func (c *getLevelCommand) Run(ctx pipeline.Context) error {
	levelId, existsUserId := ctx.GetUInt(TagLevelId)

	if !existsUserId {
		return http.CreateInternalError()
	}

	level, err := c.repository.GetLevel(levelId)

	if err != nil {
		return err
	}

	ctx.Set(TagLevel, level)
	return nil
}

func CreateGetLevelCommand(repository repository2.LevelRepository) pipeline.Step {
	return &getLevelCommand{
		repository: repository,
	}
}

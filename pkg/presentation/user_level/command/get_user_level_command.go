package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	service2 "github.com/arpb2/C-3PO/pkg/domain/user_level/service"
	levelcommand "github.com/arpb2/C-3PO/pkg/presentation/level/command"
	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"github.com/saantiaguilera/go-pipeline"
)

type getUserLevelCommand struct {
	service service2.Service
}

func (c *getUserLevelCommand) Name() string {
	return "get_user_level_command"
}

func (c *getUserLevelCommand) Run(ctx pipeline.Context) error {
	levelId, existsLevelId := ctx.GetUInt(levelcommand.TagLevelId)
	userId, existsUserId := ctx.GetUInt(command.TagUserId)

	if !existsLevelId || !existsUserId {
		return http.CreateInternalError()
	}

	userLevel, err := c.service.GetUserLevel(userId, levelId)

	if err != nil {
		return err
	}

	ctx.Set(TagUserLevel, userLevel)
	return nil
}

func CreateGetUserLevelCommand(service service2.Service) pipeline.Step {
	return &getUserLevelCommand{
		service: service,
	}
}

package user_level

import (
	"github.com/arpb2/C-3PO/api/http"
	userlevelservice "github.com/arpb2/C-3PO/api/service/user_level"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/saantiaguilera/go-pipeline"
)

type getUserLevelCommand struct {
	service userlevelservice.Service
}

func (c *getUserLevelCommand) Name() string {
	return "get_user_level_command"
}

func (c *getUserLevelCommand) Run(ctx pipeline.Context) error {
	levelId, existsLevelId := ctx.GetUInt(TagLevelId)
	userId, existsUserId := ctx.GetUInt(usercommand.TagUserId)

	if !existsLevelId || !existsUserId {
		return http.CreateInternalError()
	}

	userLevel, err := c.service.GetUserLevel(userId, levelId)

	if err != nil {
		return err
	}

	if userLevel == nil {
		return http.CreateNotFoundError()
	}

	ctx.Set(TagUserLevel, *userLevel)
	return nil
}

func CreateGetUserLevelCommand(service userlevelservice.Service) pipeline.Step {
	return &getUserLevelCommand{
		service: service,
	}
}

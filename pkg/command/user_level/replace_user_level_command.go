package user_level

import (
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	userlevelservice "github.com/arpb2/C-3PO/api/service/user_level"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/saantiaguilera/go-pipeline"
)

type replaceUserLevelCommand struct {
	service userlevelservice.Service
}

func (c *replaceUserLevelCommand) Name() string {
	return "replace_user_level_command"
}

func (c *replaceUserLevelCommand) Run(ctx pipeline.Context) error {
	levelId, existsLevelId := ctx.GetUInt(TagLevelId)
	userId, existsUserId := ctx.GetUInt(usercommand.TagUserId)
	codeRaw, existsCode := ctx.GetString(TagCodeRaw)

	if !existsLevelId || !existsUserId || !existsCode {
		return http.CreateInternalError()
	}

	userLevel := &model.UserLevel{
		LevelId: levelId,
		UserId:  userId,
		Code:    codeRaw,
	}

	err := c.service.ReplaceUserLevel(userLevel)

	if err != nil {
		return err
	}

	ctx.Set(TagUserLevel, *userLevel)
	return nil
}

func CreateReplaceUserLevelCommand(service userlevelservice.Service) pipeline.Step {
	return &replaceUserLevelCommand{
		service: service,
	}
}

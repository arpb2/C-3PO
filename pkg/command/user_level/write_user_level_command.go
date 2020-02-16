package user_level

import (
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	userlevelservice "github.com/arpb2/C-3PO/api/service/user_level"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	pipeline2 "github.com/arpb2/C-3PO/pkg/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type replaceUserLevelCommand struct {
	service userlevelservice.Service
}

func (c *replaceUserLevelCommand) Name() string {
	return "write_user_level_command"
}

func (c *replaceUserLevelCommand) Run(ctx pipeline.Context) error {
	ctxAware := pipeline2.CreateContextAware(ctx)

	levelId, existsLevelId := ctx.GetUInt(TagLevelId)
	userId, existsUserId := ctx.GetUInt(usercommand.TagUserId)
	userLevelData, existsData := ctxAware.GetUserLevelData(TagUserLevelData)

	if !existsLevelId || !existsUserId || existsData != nil {
		return http.CreateInternalError()
	}

	userLevel := model.UserLevel{
		LevelId:       levelId,
		UserId:        userId,
		UserLevelData: &userLevelData,
	}

	userLevel, err := c.service.WriteUserLevel(userLevel)

	if err != nil {
		return err
	}

	ctx.Set(TagUserLevel, userLevel)
	return nil
}

func CreateWriteUserLevelCommand(service userlevelservice.Service) pipeline.Step {
	return &replaceUserLevelCommand{
		service: service,
	}
}

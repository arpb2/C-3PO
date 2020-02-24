package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	userlevelservice "github.com/arpb2/C-3PO/pkg/domain/service/user_level"
	pipeline2 "github.com/arpb2/C-3PO/pkg/infra/pipeline"
	levelcommand "github.com/arpb2/C-3PO/pkg/presentation/level/command"
	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
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

	levelId, existsLevelId := ctx.GetUInt(levelcommand.TagLevelId)
	userId, existsUserId := ctx.GetUInt(command.TagUserId)
	userLevelData, existsData := ctxAware.GetUserLevelData(TagUserLevelData)

	if !existsLevelId || !existsUserId || existsData != nil {
		return http.CreateInternalError()
	}

	userLevel := model.UserLevel{
		LevelId:       levelId,
		UserId:        userId,
		UserLevelData: userLevelData,
	}

	userLevel, err := c.service.StoreUserLevel(userLevel)

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

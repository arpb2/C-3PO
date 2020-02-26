package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	model2 "github.com/arpb2/C-3PO/pkg/domain/user_level/model"
	repository2 "github.com/arpb2/C-3PO/pkg/domain/user_level/repository"
	levelcommand "github.com/arpb2/C-3PO/pkg/presentation/level/command"
	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"github.com/saantiaguilera/go-pipeline"
)

type replaceUserLevelCommand struct {
	repository repository2.UserLevelRepository
}

func (c *replaceUserLevelCommand) Name() string {
	return "write_user_level_command"
}

func (c *replaceUserLevelCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	levelId, existsLevelId := ctx.GetUInt(levelcommand.TagLevelId)
	userId, existsUserId := ctx.GetUInt(command.TagUserId)
	userLevelData, existsData := ctxAware.GetUserLevelData(TagUserLevelData)

	if !existsLevelId || !existsUserId || existsData != nil {
		return http.CreateInternalError()
	}

	userLevel := model2.UserLevel{
		LevelId:       levelId,
		UserId:        userId,
		UserLevelData: userLevelData,
	}

	userLevel, err := c.repository.StoreUserLevel(userLevel)

	if err != nil {
		return err
	}

	ctx.Set(TagUserLevel, userLevel)
	return nil
}

func CreateWriteUserLevelCommand(repository repository2.UserLevelRepository) pipeline.Step {
	return &replaceUserLevelCommand{
		repository: repository,
	}
}

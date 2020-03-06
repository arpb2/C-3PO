package user

import (
	user2 "github.com/arpb2/C-3PO/pkg/data/repository/user"
	level2 "github.com/arpb2/C-3PO/pkg/data/usecase/level"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type replaceUserLevelUseCase struct {
	repository user2.LevelRepository
}

func (c *replaceUserLevelUseCase) Name() string {
	return "write_user_level_usecase"
}

func (c *replaceUserLevelUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	levelId, existsLevelId := ctx.GetUInt(level2.TagLevelId)
	userId, existsUserId := ctx.GetUInt(TagUserId)
	userLevelData, existsData := ctxAware.GetUserLevelData(TagUserLevelData)

	if !existsLevelId || !existsUserId || existsData != nil {
		return http.CreateInternalError()
	}

	userLevel := user.Level{
		LevelId:   levelId,
		UserId:    userId,
		LevelData: userLevelData,
	}

	userLevel, err := c.repository.StoreUserLevel(userLevel)

	if err != nil {
		return err
	}

	ctx.Set(TagUserLevel, userLevel)
	return nil
}

func CreateWriteUserLevelUseCase(repository user2.LevelRepository) pipeline.Step {
	return &replaceUserLevelUseCase{
		repository: repository,
	}
}

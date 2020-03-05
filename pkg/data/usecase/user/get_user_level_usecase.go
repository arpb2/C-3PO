package user

import (
	user2 "github.com/arpb2/C-3PO/pkg/data/repository/user"
	level2 "github.com/arpb2/C-3PO/pkg/data/usecase/level"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/saantiaguilera/go-pipeline"
)

type getUserLevelUseCase struct {
	repository user2.LevelRepository
}

func (c *getUserLevelUseCase) Name() string {
	return "get_user_level_usecase"
}

func (c *getUserLevelUseCase) Run(ctx pipeline.Context) error {
	levelId, existsLevelId := ctx.GetUInt(level2.TagLevelId)
	userId, existsUserId := ctx.GetUInt(TagUserId)

	if !existsLevelId || !existsUserId {
		return http.CreateInternalError()
	}

	userLevel, err := c.repository.GetUserLevel(userId, levelId)

	if err != nil {
		return err
	}

	ctx.Set(TagUserLevel, userLevel)
	return nil
}

func CreateGetUserLevelUseCase(repository user2.LevelRepository) pipeline.Step {
	return &getUserLevelUseCase{
		repository: repository,
	}
}

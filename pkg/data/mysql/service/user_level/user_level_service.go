package user_level

import (
	"context"
	"time"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	userlevelservice "github.com/arpb2/C-3PO/pkg/domain/service/user_level"
	"github.com/arpb2/C-3PO/third_party/ent"
	"github.com/arpb2/C-3PO/third_party/ent/level"
	"github.com/arpb2/C-3PO/third_party/ent/user"
	"github.com/arpb2/C-3PO/third_party/ent/userlevel"
)

func CreateService(dbClient *ent.Client) userlevelservice.Service {
	return &userLevelService{
		dbClient: dbClient,
	}
}

type userLevelService struct {
	dbClient *ent.Client
}

func mapToDTO(userId, levelId uint, input *ent.UserLevel, output *model.UserLevel) {
	if input == nil {
		return
	}

	output.Code = input.Code
	output.Workspace = input.Workspace
	output.UserId = userId
	output.LevelId = levelId
}

func (c *userLevelService) GetUserLevel(userId uint, levelId uint) (userLevel model.UserLevel, err error) {
	var ul model.UserLevel
	ctx := context.Background()
	result, err := c.dbClient.UserLevel.
		Query().
		Where(
			userlevel.HasDeveloperWith(user.ID(userId)),
			userlevel.HasLevelWith(level.ID(levelId))).
		First(ctx)

	if err != nil {
		return ul, err
	}

	if result == nil {
		return ul, http.CreateNotFoundError()
	}

	mapToDTO(userId, levelId, result, &ul)

	return ul, nil
}

func (c *userLevelService) StoreUserLevel(data model.UserLevel) (userLevel model.UserLevel, err error) {
	ctx := context.Background()

	ul, err := c.dbClient.UserLevel.
		Query().
		Where(userlevel.HasDeveloperWith(user.ID(data.UserId)), userlevel.HasLevelWith(level.ID(data.LevelId))).
		First(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return data, err
	}

	if ul == nil {
		_, err = c.dbClient.UserLevel.
			Create().
			SetLevelID(data.LevelId).
			SetDeveloperID(data.UserId).
			SetCode(data.Code).
			SetWorkspace(data.Workspace).
			Save(ctx)
	} else {
		_, err = ul.
			Update().
			SetUpdatedAt(time.Now()).
			SetCode(data.Code).
			SetWorkspace(data.Workspace).
			Save(ctx)
	}

	return data, err
}

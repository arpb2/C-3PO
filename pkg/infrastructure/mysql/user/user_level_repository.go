package user

import (
	"context"
	"time"

	"github.com/arpb2/C-3PO/third_party/ent/level"

	user3 "github.com/arpb2/C-3PO/pkg/data/repository/user"
	user2 "github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/third_party/ent"
	"github.com/arpb2/C-3PO/third_party/ent/user"
	"github.com/arpb2/C-3PO/third_party/ent/userlevel"
)

func CreateLevelRepository(dbClient *ent.Client) user3.LevelRepository {
	return &userLevelRepository{
		dbClient: dbClient,
	}
}

type userLevelRepository struct {
	dbClient *ent.Client
}

func mapLevelToDTO(userId, levelId uint, input *ent.UserLevel, output *user2.Level) {
	if input == nil {
		return
	}

	output.Code = input.Code
	output.Workspace = input.Workspace
	output.UserId = userId
	output.LevelId = levelId
}

func (c *userLevelRepository) GetUserLevel(userId uint, levelId uint) (userLevel user2.Level, err error) {
	var ul user2.Level
	ctx := context.Background()
	result, err := c.dbClient.UserLevel.
		Query().
		Where(
			userlevel.HasDeveloperWith(user.ID(userId)),
			userlevel.HasLevelWith(level.ID(levelId))).
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return ul, http.CreateNotFoundError()
		}
		return ul, err
	}

	if result == nil {
		return ul, http.CreateNotFoundError()
	}

	mapLevelToDTO(userId, levelId, result, &ul)

	return ul, nil
}

func (c *userLevelRepository) GetUserLevels(userId uint) (userLevels []user2.Level, err error) {
	var uls []user2.Level
	ctx := context.Background()
	result, err := c.dbClient.UserLevel.
		Query().
		WithLevel().
		Where(userlevel.HasDeveloperWith(user.ID(userId))).
		All(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return uls, http.CreateNotFoundError()
		}
		return uls, err
	}

	if result == nil {
		return uls, http.CreateNotFoundError()
	}

	for _, ul := range result {
		var tmp user2.Level
		mapLevelToDTO(userId, ul.Edges.Level.ID, ul, &tmp)
		uls = append(uls, tmp)
	}

	return uls, nil
}

func (c *userLevelRepository) StoreUserLevel(data user2.Level) (userLevel user2.Level, err error) {
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

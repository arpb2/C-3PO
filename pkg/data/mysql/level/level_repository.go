package level

import (
	"context"
	"time"

	model2 "github.com/arpb2/C-3PO/pkg/domain/level/model"
	"github.com/arpb2/C-3PO/pkg/domain/level/repository"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/third_party/ent"
)

func CreateRepository(dbClient *ent.Client) repository.LevelRepository {
	return &levelRepository{
		dbClient: dbClient,
	}
}

type levelRepository struct {
	dbClient *ent.Client
}

func (l *levelRepository) mapToDTO(input *ent.Level, output *model2.Level) {
	if input == nil {
		return
	}

	output.Id = input.ID
	output.Name = input.Name
	output.Description = input.Description
}

func (l *levelRepository) GetLevel(levelId uint) (level model2.Level, err error) {
	var lev model2.Level
	ctx := context.Background()
	result, err := l.dbClient.Level.
		Get(ctx, levelId)

	if err != nil {
		if ent.IsNotFound(err) {
			return lev, http.CreateNotFoundError()
		}
		return lev, err
	}

	if result == nil {
		return lev, http.CreateNotFoundError()
	}

	l.mapToDTO(result, &lev)

	return lev, nil
}

func (l *levelRepository) StoreLevel(level model2.Level) (result model2.Level, err error) {
	ctx := context.Background()

	lev, err := l.dbClient.Level.
		Get(ctx, level.Id)

	if err != nil && !ent.IsNotFound(err) {
		return level, err
	}

	if lev == nil {
		_, err = l.dbClient.Level.
			Create().
			SetID(level.Id).
			SetName(level.Name).
			SetDescription(level.Description).
			Save(ctx)
	} else {
		_, err = lev.
			Update().
			SetName(level.Name).
			SetDescription(level.Description).
			SetUpdatedAt(time.Now()).
			Save(ctx)
	}

	return level, err
}

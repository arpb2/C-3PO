package level

import (
	"context"
	"time"

	level2 "github.com/arpb2/C-3PO/pkg/data/repository/level"
	"github.com/arpb2/C-3PO/pkg/domain/model/level"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/third_party/ent"
)

func CreateRepository(dbClient *ent.Client) level2.Repository {
	return &levelRepository{
		dbClient: dbClient,
	}
}

type levelRepository struct {
	dbClient *ent.Client
}

func (l *levelRepository) mapToDTO(input *ent.Level, output *level.Level) {
	if input == nil {
		return
	}

	output.Id = input.ID
	output.Name = input.Name
	output.Description = input.Description
}

func (l *levelRepository) GetLevel(levelId uint) (level.Level, error) {
	var lev level.Level
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

func (l *levelRepository) StoreLevel(lvl level.Level) (result level.Level, err error) {
	ctx := context.Background()

	lev, err := l.dbClient.Level.
		Get(ctx, lvl.Id)

	if err != nil && !ent.IsNotFound(err) {
		return lvl, err
	}

	if lev == nil {
		_, err = l.dbClient.Level.
			Create().
			SetID(lvl.Id).
			SetName(lvl.Name).
			SetDescription(lvl.Description).
			Save(ctx)
	} else {
		_, err = lev.
			Update().
			SetName(lvl.Name).
			SetDescription(lvl.Description).
			SetUpdatedAt(time.Now()).
			Save(ctx)
	}

	if err != nil && ent.IsConstraintError(err) {
		return lvl, http.CreateBadRequestError("constraint error")
	}

	return lvl, err
}

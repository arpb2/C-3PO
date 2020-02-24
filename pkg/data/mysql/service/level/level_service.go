package level

import (
	"context"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	levelservice "github.com/arpb2/C-3PO/pkg/domain/service/level"
	"github.com/arpb2/C-3PO/third_party/ent"
)

func CreateService(dbClient *ent.Client) levelservice.Service {
	return &levelService{
		dbClient: dbClient,
	}
}

type levelService struct {
	dbClient *ent.Client
}

func (l *levelService) mapToDTO(input *ent.Level, output *model.Level) {
	if input == nil {
		return
	}

	output.Id = input.ID
	output.Name = input.Name
	output.Description = input.Description
}

func (l *levelService) GetLevel(levelId uint) (level model.Level, err error) {
	var lev model.Level
	ctx := context.Background()
	result, err := l.dbClient.Level.
		Get(ctx, levelId)

	if err != nil {
		return lev, err
	}

	if result == nil {
		return lev, http.CreateNotFoundError()
	}

	l.mapToDTO(result, &lev)

	return lev, nil
}

func (l *levelService) StoreLevel(level model.Level) (result model.Level, err error) {
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
			Save(ctx)
	}

	return level, err
}

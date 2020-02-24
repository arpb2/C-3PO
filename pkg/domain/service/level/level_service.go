package level

import "github.com/arpb2/C-3PO/pkg/domain/model"

type Service interface {
	GetLevel(levelId uint) (level model.Level, err error)

	StoreLevel(level model.Level) (result model.Level, err error)
}

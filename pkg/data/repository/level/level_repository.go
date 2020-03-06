package level

import (
	"github.com/arpb2/C-3PO/pkg/domain/model/level"
)

type Repository interface {
	GetLevel(levelId uint) (level level.Level, err error)

	StoreLevel(level level.Level) (result level.Level, err error)
}

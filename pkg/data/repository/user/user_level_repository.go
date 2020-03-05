package user

import (
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
)

type LevelRepository interface {
	GetUserLevel(userId uint, levelId uint) (userLevel user.Level, err error)

	StoreUserLevel(userLevel user.Level) (model user.Level, err error)
}

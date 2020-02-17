package user_level

import "github.com/arpb2/C-3PO/api/model"

type Service interface {
	GetUserLevel(userId uint, levelId uint) (userLevel model.UserLevel, err error)

	StoreUserLevel(userLevel model.UserLevel) (model model.UserLevel, err error)
}

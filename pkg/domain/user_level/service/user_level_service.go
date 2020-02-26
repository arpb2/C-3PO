package service

import (
	model2 "github.com/arpb2/C-3PO/pkg/domain/user_level/model"
)

type Service interface {
	GetUserLevel(userId uint, levelId uint) (userLevel model2.UserLevel, err error)

	StoreUserLevel(userLevel model2.UserLevel) (model model2.UserLevel, err error)
}

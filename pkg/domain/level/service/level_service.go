package service

import (
	model2 "github.com/arpb2/C-3PO/pkg/domain/level/model"
)

type Service interface {
	GetLevel(levelId uint) (level model2.Level, err error)

	StoreLevel(level model2.Level) (result model2.Level, err error)
}

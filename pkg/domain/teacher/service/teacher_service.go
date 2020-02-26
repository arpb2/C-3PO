package service

import (
	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
)

type Service interface {
	GetStudents(userId uint) (students []model2.User, err error)
}

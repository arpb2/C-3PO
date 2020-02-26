package repository

import (
	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
)

type TeacherRepository interface {
	GetStudents(userId uint) (students []model2.User, err error)
}

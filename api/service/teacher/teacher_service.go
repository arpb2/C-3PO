package teacher_service

import "github.com/arpb2/C-3PO/api/model"

type Service interface {
	GetStudents(userId uint) (students *[]model.User, err error)
}

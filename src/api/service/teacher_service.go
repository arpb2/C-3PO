package service

import "github.com/arpb2/C-3PO/src/api/model"

type TeacherService interface {

	GetStudents(userId uint) (students *[]model.User, err error)

}

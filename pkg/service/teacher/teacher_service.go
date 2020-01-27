package teacher_service

import (
	"github.com/arpb2/C-3PO/api/model"
	teacher_service "github.com/arpb2/C-3PO/api/service/teacher"
	user_service "github.com/arpb2/C-3PO/api/service/user"
)

func CreateService(userService user_service.Service) teacher_service.Service {
	return &teacherService{
		Service: userService,
	}
}

type teacherService struct {
	user_service.Service
}

func (t teacherService) GetStudents(userId uint) (students *[]model.User, err error) {
	panic("implement me")
}

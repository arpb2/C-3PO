package teacher

import (
	"github.com/arpb2/C-3PO/api/model"
	teacherservice "github.com/arpb2/C-3PO/api/service/teacher"
	userservice "github.com/arpb2/C-3PO/api/service/user"
)

func CreateService(userService userservice.Service) teacherservice.Service {
	return &teacherService{
		Service: userService,
	}
}

type teacherService struct {
	userservice.Service
}

func (t teacherService) GetStudents(userId uint) (students []model.User, err error) {
	panic("implement me")
}

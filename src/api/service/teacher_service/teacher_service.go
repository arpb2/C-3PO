package teacher_service

import (
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

func CreateService(userService service.UserService) service.TeacherService {
	return &teacherService{
		UserService: userService,
	}
}

type teacherService struct {
	service.UserService
}
func (t teacherService) GetStudents(userId uint) (students *[]model.User, err error) {
	panic("implement me")
}




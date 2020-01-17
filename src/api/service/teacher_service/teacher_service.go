package teacher_service

import (
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/arpb2/C-3PO/src/api/service/user_service"
)

func GetService() service.TeacherService {
	return globalRef
}

var globalRef = &teacherService{
	UserService: user_service.GetService(),
}

type teacherService struct {
	service.UserService
}
func (t teacherService) GetStudents(userId uint) (students *[]model.User, err error) {
	panic("implement me")
}




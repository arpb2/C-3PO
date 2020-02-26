package teacher

import (
	service2 "github.com/arpb2/C-3PO/pkg/domain/teacher/service"
	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
	"github.com/arpb2/C-3PO/pkg/domain/user/service"
	"github.com/arpb2/C-3PO/third_party/ent"
)

func CreateService(userService service.Service, dbClient *ent.Client) service2.Service {
	return &teacherService{
		Service:  userService,
		dbClient: dbClient,
	}
}

type teacherService struct {
	service.Service
	dbClient *ent.Client
}

func (t teacherService) GetStudents(userId uint) (students []model2.User, err error) {
	panic("implement me")
}

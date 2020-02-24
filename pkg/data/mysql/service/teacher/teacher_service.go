package teacher

import (
	"github.com/arpb2/C-3PO/pkg/domain/model"
	teacherservice "github.com/arpb2/C-3PO/pkg/domain/service/teacher"
	userservice "github.com/arpb2/C-3PO/pkg/domain/service/user"
	"github.com/arpb2/C-3PO/third_party/ent"
)

func CreateService(userService userservice.Service, dbClient *ent.Client) teacherservice.Service {
	return &teacherService{
		Service:  userService,
		dbClient: dbClient,
	}
}

type teacherService struct {
	userservice.Service
	dbClient *ent.Client
}

func (t teacherService) GetStudents(userId uint) (students []model.User, err error) {
	panic("implement me")
}

package teacher

import (
	"github.com/arpb2/C-3PO/api/model"
	teacherservice "github.com/arpb2/C-3PO/api/service/teacher"
	userservice "github.com/arpb2/C-3PO/api/service/user"
	"github.com/arpb2/C-3PO/pkg/ent"
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

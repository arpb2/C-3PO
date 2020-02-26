package teacher

import (
	repository2 "github.com/arpb2/C-3PO/pkg/domain/teacher/repository"
	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
	"github.com/arpb2/C-3PO/pkg/domain/user/repository"
	"github.com/arpb2/C-3PO/third_party/ent"
)

func CreateRepository(userRepository repository.UserRepository, dbClient *ent.Client) repository2.TeacherRepository {
	return &teacherRepository{
		UserRepository: userRepository,
		dbClient:       dbClient,
	}
}

type teacherRepository struct {
	repository.UserRepository
	dbClient *ent.Client
}

func (t teacherRepository) GetStudents(userId uint) (students []model2.User, err error) {
	return []model2.User{}, nil
}

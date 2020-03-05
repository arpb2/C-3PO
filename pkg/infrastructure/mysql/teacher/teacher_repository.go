package teacher

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/teacher"
	user2 "github.com/arpb2/C-3PO/pkg/data/repository/user"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
	"github.com/arpb2/C-3PO/third_party/ent"
)

func CreateRepository(userRepository user2.Repository, dbClient *ent.Client) teacher.StudentRepository {
	return &teacherRepository{
		Repository: userRepository,
		dbClient:   dbClient,
	}
}

type teacherRepository struct {
	user2.Repository
	dbClient *ent.Client
}

func (t teacherRepository) GetStudents(userId uint) (students []user.User, err error) {
	return []user.User{}, nil
}

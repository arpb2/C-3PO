package teacher

import (
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
)

type StudentRepository interface {
	GetStudents(userId uint) (students []user.User, err error)
}

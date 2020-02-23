package teacher

import "github.com/arpb2/C-3PO/pkg/domain/model"

type Service interface {
	GetStudents(userId uint) (students []model.User, err error)
}

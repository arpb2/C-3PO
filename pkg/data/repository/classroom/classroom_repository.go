package classroom

import (
	"github.com/arpb2/C-3PO/pkg/domain/model/classroom"
)

type Repository interface {
	GetClassroom(classroomID uint) (classroom classroom.Classroom, err error)

	UpdateClassroom(classroom classroom.Classroom) (result classroom.Classroom, err error)
}

package classroom

import (
	"context"
	"github.com/arpb2/C-3PO/pkg/data/repository/classroom"
	classroom2 "github.com/arpb2/C-3PO/pkg/domain/model/classroom"
	user2 "github.com/arpb2/C-3PO/pkg/domain/model/user"
	classroom3 "github.com/arpb2/C-3PO/third_party/ent/classroom"
	"github.com/arpb2/C-3PO/third_party/ent/user"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/third_party/ent"
)

func CreateRepository(dbClient *ent.Client) classroom.Repository {
	return &classroomRepository{
		dbClient: dbClient,
	}
}

type classroomRepository struct {
	dbClient *ent.Client
}

func mapToClassroom(cr *ent.Classroom) classroom2.Classroom {
	var students []user2.User
	for _, st := range cr.Edges.Students {
		students = append(students, user2.User{
			Id:          st.ID,
			Type:        user2.Type(st.Type),
			ClassroomID: cr.ID,
			Email:       st.Email,
			Name:        st.Name,
			Surname:     st.Surname,
		})
	}
	return classroom2.Classroom{
		Id:       cr.ID,
		Level:    cr.Edges.Level.ID,
		Students: students,
		Teacher:  user2.User{
			Id:          cr.Edges.Teacher.ID,
			Type:        user2.Type(cr.Edges.Teacher.Type),
			ClassroomID: cr.ID,
			Email:       cr.Edges.Teacher.Email,
			Name:        cr.Edges.Teacher.Name,
			Surname:     cr.Edges.Teacher.Surname,
		},
	}
}

func (c classroomRepository) GetClassroom(classroomID uint) (classroom classroom2.Classroom, err error) {
	ctx := context.Background()
	cr, err := c.dbClient.Classroom.
		Query().
		WithLevel().
		WithStudents().
		WithTeacher().
		Where(classroom3.ID(classroomID)).
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return classroom2.Classroom{}, http.CreateUnauthorizedError()
		}
		return classroom2.Classroom{}, err
	}

	if cr == nil {
		return classroom2.Classroom{}, http.CreateInternalError()
	}

	return mapToClassroom(cr), nil
}

func (c classroomRepository) UpdateClassroom(classroom classroom2.Classroom) (result classroom2.Classroom, err error) {
	ctx := context.Background()
	update := c.dbClient.Classroom.UpdateOneID(classroom.Id)

	if classroom.Level != 0 {
		update.SetLevelID(classroom.Level)
	}

	if classroom.Students != nil && len(classroom.Students) > 0 {
		var emails []string
		for _, st := range classroom.Students {
			emails = append(emails, st.Email)
		}
		students, err := c.dbClient.User.
			Query().
			Where(user.EmailIn(emails...)).
			All(ctx)
		if err != nil {
			return classroom2.Classroom{}, err
		}

		var ids []uint
		for _, st := range students {
			if st.Type == user.TypeTeacher {
				return classroom2.Classroom{}, http.CreateBadRequestError("a student is of type teacher")
			}
			ids = append(ids, st.ID)
		}
		r, err := c.dbClient.Classroom.Query().Where(classroom3.HasStudentsWith(user.IDIn(ids...))).All(ctx)

		if err != nil {
			return classroom2.Classroom{}, err
		}
		if len(r) > 0 {
			return classroom2.Classroom{}, http.CreateBadRequestError("some students already have a classroom associated")
		}

		update.AddStudentIDs(ids...)
	}

	cr, err := update.Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return classroom2.Classroom{}, http.CreateUnauthorizedError()
		}
		return classroom2.Classroom{}, err
	}

	if cr == nil {
		return classroom2.Classroom{}, http.CreateInternalError()
	}

	return mapToClassroom(cr), nil
}

package teacher

import (
	"strconv"

	sessionrepository "github.com/arpb2/C-3PO/pkg/domain/session/repository"

	teacherrepository "github.com/arpb2/C-3PO/pkg/domain/teacher/repository"

	"github.com/arpb2/C-3PO/pkg/presentation/middleware/auth/user"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
)

func CreateMiddleware(tokenHandler sessionrepository.TokenRepository, teacherRepository teacherrepository.TeacherRepository) http.Handler {
	strategy := &teacherAuthenticationStrategy{
		teacherRepository,
	}

	return func(ctx *http.Context) {
		user.HandleAuthentication(ctx, tokenHandler, strategy)
	}
}

type teacherAuthenticationStrategy struct {
	teacherrepository.TeacherRepository
}

func (s teacherAuthenticationStrategy) Authenticate(token *sessionrepository.Token, userId string) (bool, error) {
	students, err := s.TeacherRepository.GetStudents(token.UserId)

	if err != nil {
		return false, err
	}

	if students != nil {
		for _, student := range students {
			if strconv.FormatUint(uint64(student.Id), 10) == userId {
				return true, nil
			}
		}
	}

	return false, nil
}

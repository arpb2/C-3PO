package session

import (
	"strconv"

	"github.com/arpb2/C-3PO/pkg/data/repository/session"
	"github.com/arpb2/C-3PO/pkg/data/repository/teacher"
)

func CreateTeacherAuthenticationUseCase(
	tokenHandler session.TokenRepository,
	teacherRepository teacher.StudentRepository,
) func(string, string) error {
	strategy := &teacherAuthenticationStrategy{
		teacherRepository,
	}

	return func(authToken, userId string) error {
		return HandleTokenizedAuthentication(authToken, userId, tokenHandler, strategy)
	}
}

type teacherAuthenticationStrategy struct {
	teacher.StudentRepository
}

func (s teacherAuthenticationStrategy) Authenticate(token *session.Token, userId string) (bool, error) {
	students, err := s.StudentRepository.GetStudents(token.UserId)

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

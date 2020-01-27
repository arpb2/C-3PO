package teacher_auth

import (
	"strconv"

	"github.com/arpb2/C-3PO/api/auth"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	teacher_service "github.com/arpb2/C-3PO/api/service/teacher"
	middleware_auth "github.com/arpb2/C-3PO/pkg/middleware/auth"
)

func CreateMiddleware(tokenHandler auth.TokenHandler, teacherService teacher_service.Service) http_wrapper.Handler {
	strategy := &teacherAuthenticationStrategy{
		teacherService,
	}

	return func(ctx *http_wrapper.Context) {
		middleware_auth.HandleAuthentication(ctx, tokenHandler, strategy)
	}
}

type teacherAuthenticationStrategy struct {
	teacher_service.Service
}

func (s teacherAuthenticationStrategy) Authenticate(token *auth.Token, userId string) (bool, error) {
	students, err := s.Service.GetStudents(token.UserId)

	if err != nil {
		return false, err
	}

	if students != nil {
		for _, student := range *students {
			if strconv.FormatUint(uint64(student.Id), 10) == userId {
				return true, nil
			}
		}
	}

	return false, nil
}

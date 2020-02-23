package teacher

import (
	"strconv"

	"github.com/arpb2/C-3PO/pkg/presentation/auth/middleware"

	"github.com/arpb2/C-3PO/pkg/domain/auth"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	teacherservice "github.com/arpb2/C-3PO/pkg/domain/service/teacher"
)

func CreateMiddleware(tokenHandler auth.TokenHandler, teacherService teacherservice.Service) http.Handler {
	strategy := &teacherAuthenticationStrategy{
		teacherService,
	}

	return func(ctx *http.Context) {
		middleware.HandleAuthentication(ctx, tokenHandler, strategy)
	}
}

type teacherAuthenticationStrategy struct {
	teacherservice.Service
}

func (s teacherAuthenticationStrategy) Authenticate(token *auth.Token, userId string) (bool, error) {
	students, err := s.Service.GetStudents(token.UserId)

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

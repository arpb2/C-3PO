package teacher

import (
	"strconv"

	"github.com/arpb2/C-3PO/api/auth"
	"github.com/arpb2/C-3PO/api/http"
	teacherservice "github.com/arpb2/C-3PO/api/service/teacher"
	middlewareauth "github.com/arpb2/C-3PO/pkg/middleware/auth"
)

func CreateMiddleware(tokenHandler auth.TokenHandler, teacherService teacherservice.Service) http.Handler {
	strategy := &teacherAuthenticationStrategy{
		teacherService,
	}

	return func(ctx *http.Context) {
		middlewareauth.HandleAuthentication(ctx, tokenHandler, strategy)
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
		for _, student := range *students {
			if strconv.FormatUint(uint64(student.Id), 10) == userId {
				return true, nil
			}
		}
	}

	return false, nil
}

package teacher

import (
	"strconv"

	"github.com/arpb2/C-3PO/pkg/domain/session/token"
	"github.com/arpb2/C-3PO/pkg/domain/teacher/service"

	"github.com/arpb2/C-3PO/pkg/presentation/middleware/user"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
)

func CreateMiddleware(tokenHandler token.Handler, teacherService service.Service) http.Handler {
	strategy := &teacherAuthenticationStrategy{
		teacherService,
	}

	return func(ctx *http.Context) {
		user.HandleAuthentication(ctx, tokenHandler, strategy)
	}
}

type teacherAuthenticationStrategy struct {
	service.Service
}

func (s teacherAuthenticationStrategy) Authenticate(token *token.Token, userId string) (bool, error) {
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

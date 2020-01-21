package teacher_auth

import (
	"github.com/arpb2/C-3PO/src/api/auth"
	middleware_auth "github.com/arpb2/C-3PO/src/api/middleware/auth"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateMiddleware(tokenHandler auth.TokenHandler, teacherService service.TeacherService) gin.HandlerFunc {
	strategy := &teacherAuthenticationStrategy{
		teacherService,
	}

	return func(ctx *gin.Context) {
		middleware_auth.HandleAuthentication(ctx, tokenHandler, strategy)
	}
}

type teacherAuthenticationStrategy struct {
	service.TeacherService
}

func (s teacherAuthenticationStrategy) Authenticate(token *auth.Token, userId string) (bool, error) {
	students, err := s.TeacherService.GetStudents(token.UserId)

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

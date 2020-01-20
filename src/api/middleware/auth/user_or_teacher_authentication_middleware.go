package auth

import (
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

var TeacherService service.TeacherService // TODO Set.

var UserOrTeacherAuthenticationMiddleware gin.HandlerFunc = userOrTeacherAuthenticationProxy

func userOrTeacherAuthenticationProxy(ctx *gin.Context) {
	handleAuthentication(ctx, teacherAuthenticationStrategy)
}

func teacherAuthenticationStrategy(token *auth.Token, userId string) (bool, error) {
	students, err := TeacherService.GetStudents(token.UserId)

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

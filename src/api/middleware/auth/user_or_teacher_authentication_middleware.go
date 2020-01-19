package auth

import "github.com/gin-gonic/gin"

var UserOrTeacherAuthenticationMiddleware gin.HandlerFunc = userOrTeacherAuthenticationProxy

func userOrTeacherAuthenticationProxy(ctx *gin.Context) {
	handleAuthentication(ctx, true)
}

package auth

import "github.com/gin-gonic/gin"

var SingleAuthenticationMiddleware gin.HandlerFunc = singleAuthenticationProxy

func singleAuthenticationProxy(ctx *gin.Context) {
	handleAuthentication(ctx)
}
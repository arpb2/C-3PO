package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Method string

	Path string

	Middleware []gin.HandlerFunc

	Body gin.HandlerFunc
}

func Halt(ctx *gin.Context, code int, errMessage string) {
	if code >= 200 && code < 300 {
		fmt.Printf(
			"Request from %s was requested to be halted with code '%d' and message '%s' when its a successful repsonse",
			ctx.Request.URL.String(),
			code,
			errMessage,
		)
	} else {
		ctx.AbortWithStatusJSON(code, gin.H{
			"error": errMessage,
		})
	}
}
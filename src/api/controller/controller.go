package controller

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
)

type Controller struct {
	Method string

	Path string

	Middleware []http_wrapper.Handler

	Body http_wrapper.Handler
}

func Halt(ctx *http_wrapper.Context, code int, errMessage string) {
	if code >= 200 && code < 300 {
		fmt.Printf(
			"Request from %s was requested to be halted with code '%d' and message '%s' when its a successful response",
			ctx.Url(),
			code,
			errMessage,
		)
	} else {
		ctx.AbortWithStatusJSON(code, http_wrapper.Json{
			"error": errMessage,
		})
	}
}
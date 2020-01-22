package controller

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"golang.org/x/xerrors"
)

type Controller struct {
	Method string

	Path string

	Middleware []http_wrapper.Handler

	Body http_wrapper.Handler
}

func HaltError(ctx *http_wrapper.Context, err error) error {
	var httpError http_wrapper.HttpError
	if xerrors.As(err, &httpError) {
		Halt(ctx, httpError.Code, httpError.Error())
		return nil
	}

	return err
}

func Halt(ctx *http_wrapper.Context, code int, errMessage string) {
	if code >= 200 && code < 300 {
		fmt.Printf(
			"Request from %s was requested to be halted with code '%d' and message '%s' when its a successful response",
			ctx.GetUrl(),
			code,
			errMessage,
		)
	} else {
		ctx.AbortTransactionWithStatus(code, http_wrapper.Json{
			"error": errMessage,
		})
	}
}
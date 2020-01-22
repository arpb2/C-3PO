package controller

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"golang.org/x/xerrors"
	"net/http"
)

type Controller struct {
	Method string

	Path string

	Middleware []http_wrapper.Handler

	Body http_wrapper.Handler
}

func HaltExternalError(ctx *http_wrapper.Context, err error) error {
	var httpError http_wrapper.HttpError
	if xerrors.As(err, &httpError) && httpError.Code != http.StatusInternalServerError {
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

func BatchRun(exec executor.Executor, commands []executor.Command, ctx *http_wrapper.Context) error {
	var channels []<-chan error
	for _, command := range commands {
		errChan := exec.Go(command)

		channels = append(channels, errChan)
	}

	singleErrorChannel := executor.Merge(channels...)

	if err, open := <-singleErrorChannel; open {
		fmt.Print(err.Error())
		Halt(ctx, http.StatusInternalServerError, "internal error")
		return err
	}
	return nil
}
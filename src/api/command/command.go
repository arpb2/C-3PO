package command

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"golang.org/x/xerrors"
	"net/http"
)

func HaltClientHttpError(ctx *http_wrapper.Context, err error) error {
	var httpError http_wrapper.HttpError
	if xerrors.As(err, &httpError) && httpError.Code < http.StatusInternalServerError {
		ctx.AbortTransactionWithError(httpError)
		return nil
	}

	return err
}


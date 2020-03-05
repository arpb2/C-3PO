package health

import (
	"net/http"

	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
)

func CreateGetHandler() httpwrapper.Handler {
	return func(ctx *httpwrapper.Context) {
		ctx.WriteString(http.StatusOK, "pong")
	}
}

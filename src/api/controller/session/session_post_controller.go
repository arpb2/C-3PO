package session

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"net/http"
)

func CreatePostController() controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/session",
		Body:   CreatePostBody(),
	}
}

func CreatePostBody() http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		ctx.WriteString(http.StatusOK, "pong")
	}
}

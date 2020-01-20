package user

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/controller"
	"net/http"
)

func CreatePostController(authMiddleware http_wrapper.Handler) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/users",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body: userPost,
	}
}

func userPost(ctx *http_wrapper.Context) {
	ctx.JSON(http.StatusOK, http_wrapper.Json{})
}

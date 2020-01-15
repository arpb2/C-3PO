package user

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/middleware/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

var DeleteController = controller.Controller{
	Method: "DELETE",
	Path:   "/users/:user_id",
	Middleware: []gin.HandlerFunc{
		auth.SingleAuthenticationMiddleware,
	},
	Body:   userDelete,
}

func userDelete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
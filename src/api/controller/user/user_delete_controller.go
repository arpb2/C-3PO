package user

import (
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/single_auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateDeleteController() controller.Controller {
	return controller.Controller{
		Method: "DELETE",
		Path:   "/users/:user_id",
		Middleware: []gin.HandlerFunc{
			single_auth.CreateMiddleware(
				jwt.CreateTokenHandler(),
			),
		},
		Body:   userDelete,
	}
}

func userDelete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
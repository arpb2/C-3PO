package user

import (
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/single_auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGetController() controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   "/users/:user_id",
		Middleware: []gin.HandlerFunc{
			single_auth.CreateMiddleware(
				jwt.CreateTokenHandler(),
			),
		},
		Body:   userGet,
	}
}

func userGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
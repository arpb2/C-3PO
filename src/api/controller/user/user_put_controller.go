package user

import (
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/single_auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePutController() controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   "/users/:user_id",
		Middleware: []gin.HandlerFunc{
			single_auth.CreateMiddleware(
				jwt.CreateTokenHandler(),
			),
		},
		Body:   userPut,
	}
}

func userPut(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

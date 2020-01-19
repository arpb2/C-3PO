package user

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

var PutController = controller.Controller{
	Method: "PUT",
	Path:   "/users/:user_id",
	Body:   userPut,
}

func userPut(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

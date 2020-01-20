package user_test

import (
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/single_auth"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createGetController() controller.Controller {
	return user.CreateGetController(
		single_auth.CreateMiddleware(
			jwt.CreateTokenHandler(),
		),
	)
}

func TestUserGetControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", createGetController().Method)
}

func TestUserGetControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id", createGetController().Path)
}
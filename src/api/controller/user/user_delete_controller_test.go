package user_test

import (
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/single_auth"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createDeleteController() controller.Controller {
	return user.CreateDeleteController(
		single_auth.CreateMiddleware(
			jwt.CreateTokenHandler(),
		),
	)
}

func TestUserDeleteControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "DELETE", createDeleteController().Method)
}

func TestUserDeleteControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id", createDeleteController().Path)
}
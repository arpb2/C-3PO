package user_test

import (
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/single_auth"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createPutController() controller.Controller {
	return user.CreatePutController(
		single_auth.CreateMiddleware(
			jwt.CreateTokenHandler(),
		),
	)
}

func TestUserPutControllerMethodIsPUT(t *testing.T) {
	assert.Equal(t, "PUT", createPutController().Method)
}

func TestUserPutControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id", createPutController().Path)
}

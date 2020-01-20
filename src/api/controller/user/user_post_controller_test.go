package user_test

import (
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/single_auth"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createPostController() controller.Controller {
	return user.CreatePostController(
		single_auth.CreateMiddleware(
			jwt.CreateTokenHandler(),
		),
	)
}

func TestUserPostControllerMethodIsPOST(t *testing.T) {
	assert.Equal(t, "POST", createPostController().Method)
}

func TestUserPostControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users", createPostController().Path)
}

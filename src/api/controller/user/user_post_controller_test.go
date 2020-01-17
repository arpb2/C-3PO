package user_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/arpb2/C-3PO/src/api/middleware/auth"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestUserPostControllerMethodIsPOST(t *testing.T) {
	assert.Equal(t, "POST", user.PostController.Method)
}

func TestUserPostControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users", user.PostController.Path)
}

func TestUserPostControllerMiddleware_HasAuthenticationMiddleware(t *testing.T) {
	found := false

	for _, middleware := range user.PostController.Middleware {
		// Golang doesn't allow func comparisons, so we have to test identity through pointers using reflection.
		if reflect.ValueOf(auth.SingleAuthenticationMiddleware).Pointer() == reflect.ValueOf(middleware).Pointer() {
			found = true
		}
	}

	assert.True(t, found)
}
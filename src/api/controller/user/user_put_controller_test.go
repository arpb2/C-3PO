package user_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/arpb2/C-3PO/src/api/middleware/auth"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestUserPutControllerMethodIsPUT(t *testing.T) {
	assert.Equal(t, "PUT", user.PutController.Method)
}

func TestUserPutControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id", user.PutController.Path)
}

func TestUserPutControllerMiddleware_HasAuthenticationMiddleware(t *testing.T) {
	found := false

	for _, middleware := range user.PutController.Middleware {
		// Golang doesn't allow func comparisons, so we have to test identity through pointers using reflection.
		if reflect.ValueOf(auth.SingleAuthenticationMiddleware).Pointer() == reflect.ValueOf(middleware).Pointer() {
			found = true
		}
	}

	assert.True(t, found)
}

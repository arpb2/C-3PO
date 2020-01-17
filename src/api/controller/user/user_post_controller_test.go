package user_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserPostControllerMethodIsPOST(t *testing.T) {
	assert.Equal(t, "POST", user.CreatePostController().Method)
}

func TestUserPostControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users", user.CreatePostController().Path)
}

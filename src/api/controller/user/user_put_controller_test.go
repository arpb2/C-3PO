package user_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserPutControllerMethodIsPUT(t *testing.T) {
	assert.Equal(t, "PUT", user.PutController.Method)
}

func TestUserPutControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id", user.PutController.Path)
}

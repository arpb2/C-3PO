package user_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserGetControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", user.CreateGetController().Method)
}

func TestUserGetControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id", user.CreateGetController().Path)
}
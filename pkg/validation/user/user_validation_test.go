package user_validation_test

import (
	"testing"

	"github.com/arpb2/C-3PO/api/model"
	user_validation "github.com/arpb2/C-3PO/pkg/validation/user"
	"github.com/stretchr/testify/assert"
)

func TestUserValidation_Error(t *testing.T) {
	err := user_validation.EmptyUser(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "no user found", err.Error())
}

func TestUserValidation_Success(t *testing.T) {
	err := user_validation.EmptyUser(&model.AuthenticatedUser{})

	assert.Nil(t, err)
}

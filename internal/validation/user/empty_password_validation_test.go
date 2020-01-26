package user_validation_test

import (
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/internal/validation/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyPasswordValidation_Error(t *testing.T) {
	err := user_validation.EmptyPassword(&model.AuthenticatedUser{
		User:     &model.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'password' provided", err.Error())
}

func TestEmptyPasswordValidation_Success(t *testing.T) {
	err := user_validation.EmptyPassword(&model.AuthenticatedUser{
		User:     &model.User{},
		Password: "test password",
	})

	assert.Nil(t, err)
}

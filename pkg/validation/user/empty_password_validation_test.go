package user_validation_test

import (
	"testing"

	"github.com/arpb2/C-3PO/api/model"
	user_validation "github.com/arpb2/C-3PO/pkg/validation/user"
	"github.com/stretchr/testify/assert"
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

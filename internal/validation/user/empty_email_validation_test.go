package user_validation_test

import (
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/internal/validation/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyEmailValidation_Error(t *testing.T) {
	err := user_validation.EmptyEmail(&model.AuthenticatedUser{
		User:     &model.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'email' provided", err.Error())
}

func TestEmptyEmailValidation_Success(t *testing.T) {
	err := user_validation.EmptyEmail(&model.AuthenticatedUser{
		User: &model.User{
			Email: "test@email.com",
		},
		Password: "",
	})

	assert.Nil(t, err)
}

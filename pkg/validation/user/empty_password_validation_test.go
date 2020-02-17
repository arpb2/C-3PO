package user_test

import (
	"testing"

	"github.com/arpb2/C-3PO/api/model"
	uservalidation "github.com/arpb2/C-3PO/pkg/validation/user"
	"github.com/stretchr/testify/assert"
)

func TestEmptyPasswordValidation_Error(t *testing.T) {
	err := uservalidation.EmptyPassword(&model.AuthenticatedUser{
		User:     model.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'password' provided", err.Error())
}

func TestEmptyPasswordValidation_Success(t *testing.T) {
	err := uservalidation.EmptyPassword(&model.AuthenticatedUser{
		User:     model.User{},
		Password: "test password",
	})

	assert.Nil(t, err)
}

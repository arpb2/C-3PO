package validation_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user/validation"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestEmptyPasswordValidation_Error(t *testing.T) {
	err := validation.EmptyPassword(&user.AuthenticatedUser{
		User:     user.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'password' provided", err.Error())
}

func TestEmptyPasswordValidation_Success(t *testing.T) {
	err := validation.EmptyPassword(&user.AuthenticatedUser{
		User:     user.User{},
		Password: "test password",
	})

	assert.Nil(t, err)
}

package validation_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user/validation"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestEmptyEmailValidation_Error(t *testing.T) {
	err := validation.EmptyEmail(&user.AuthenticatedUser{
		User:     user.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'email' provided", err.Error())
}

func TestEmptyEmailValidation_Success(t *testing.T) {
	err := validation.EmptyEmail(&user.AuthenticatedUser{
		User: user.User{
			Email: "test@email.com",
		},
		Password: "",
	})

	assert.Nil(t, err)
}

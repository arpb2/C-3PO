package validation_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user/validation"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestEmptyNameValidation_Error(t *testing.T) {
	err := validation.EmptyName(&user.AuthenticatedUser{
		User:     user.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'name' provided", err.Error())
}

func TestEmptyNameValidation_Success(t *testing.T) {
	err := validation.EmptyName(&user.AuthenticatedUser{
		User: user.User{
			Name: "TestName",
		},
		Password: "",
	})

	assert.Nil(t, err)
}

package validation_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestEmptyPasswordValidation_Error(t *testing.T) {
	err := validation.EmptyPassword(&model.AuthenticatedUser{
		User:     model.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'password' provided", err.Error())
}

func TestEmptyPasswordValidation_Success(t *testing.T) {
	err := validation.EmptyPassword(&model.AuthenticatedUser{
		User:     model.User{},
		Password: "test password",
	})

	assert.Nil(t, err)
}

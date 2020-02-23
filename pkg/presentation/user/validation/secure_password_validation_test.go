package validation_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestSecurePasswordValidation_Error(t *testing.T) {
	err := validation.SecurePassword(&model.AuthenticatedUser{
		User:     model.User{},
		Password: "1234567",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "password must have at least 8 characters", err.Error())
}

func TestSecurePasswordValidation_Success(t *testing.T) {
	err := validation.SecurePassword(&model.AuthenticatedUser{
		User:     model.User{},
		Password: "12345678",
	})

	assert.Nil(t, err)
}

package validation_test

import (
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"

	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/stretchr/testify/assert"
)

func TestSecurePasswordValidation_Error(t *testing.T) {
	err := validation.SecurePassword(&model2.AuthenticatedUser{
		User:     model2.User{},
		Password: "1234567",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "password must have at least 8 characters", err.Error())
}

func TestSecurePasswordValidation_Success(t *testing.T) {
	err := validation.SecurePassword(&model2.AuthenticatedUser{
		User:     model2.User{},
		Password: "12345678",
	})

	assert.Nil(t, err)
}

package validation_test

import (
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"

	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/stretchr/testify/assert"
)

func TestEmptyPasswordValidation_Error(t *testing.T) {
	err := validation.EmptyPassword(&model2.AuthenticatedUser{
		User:     model2.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'password' provided", err.Error())
}

func TestEmptyPasswordValidation_Success(t *testing.T) {
	err := validation.EmptyPassword(&model2.AuthenticatedUser{
		User:     model2.User{},
		Password: "test password",
	})

	assert.Nil(t, err)
}

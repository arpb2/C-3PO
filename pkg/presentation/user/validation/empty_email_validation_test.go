package validation_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestEmptyEmailValidation_Error(t *testing.T) {
	err := validation.EmptyEmail(&model.AuthenticatedUser{
		User:     model.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'email' provided", err.Error())
}

func TestEmptyEmailValidation_Success(t *testing.T) {
	err := validation.EmptyEmail(&model.AuthenticatedUser{
		User: model.User{
			Email: "test@email.com",
		},
		Password: "",
	})

	assert.Nil(t, err)
}

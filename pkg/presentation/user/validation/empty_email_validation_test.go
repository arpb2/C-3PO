package validation_test

import (
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"

	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/stretchr/testify/assert"
)

func TestEmptyEmailValidation_Error(t *testing.T) {
	err := validation.EmptyEmail(&model2.AuthenticatedUser{
		User:     model2.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'email' provided", err.Error())
}

func TestEmptyEmailValidation_Success(t *testing.T) {
	err := validation.EmptyEmail(&model2.AuthenticatedUser{
		User: model2.User{
			Email: "test@email.com",
		},
		Password: "",
	})

	assert.Nil(t, err)
}

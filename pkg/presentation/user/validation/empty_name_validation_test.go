package validation_test

import (
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"

	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/stretchr/testify/assert"
)

func TestEmptyNameValidation_Error(t *testing.T) {
	err := validation.EmptyName(&model2.AuthenticatedUser{
		User:     model2.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'name' provided", err.Error())
}

func TestEmptyNameValidation_Success(t *testing.T) {
	err := validation.EmptyName(&model2.AuthenticatedUser{
		User: model2.User{
			Name: "TestName",
		},
		Password: "",
	})

	assert.Nil(t, err)
}

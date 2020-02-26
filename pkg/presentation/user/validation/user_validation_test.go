package validation_test

import (
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"

	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/stretchr/testify/assert"
)

func TestUserValidation_Error(t *testing.T) {
	err := validation.EmptyUser(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "no user found", err.Error())
}

func TestUserValidation_Success(t *testing.T) {
	err := validation.EmptyUser(&model2.AuthenticatedUser{})

	assert.Nil(t, err)
}

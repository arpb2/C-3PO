package validation_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user/validation"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestMalformedTypeValidation_Error(t *testing.T) {
	err := validation.MalformedType(&user.AuthenticatedUser{
		User: user.User{
			Type: "some type",
		},
	})

	assert.NotNil(t, err)
	assert.Equal(t, "'type' is malformed", err.Error())
}

func TestMalformedTypeValidation_Success(t *testing.T) {
	err := validation.MalformedType(&user.AuthenticatedUser{
		User: user.User{
			Type: user.TypeStudent,
		},
	})

	assert.Nil(t, err)
}

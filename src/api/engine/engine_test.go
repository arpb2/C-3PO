package engine_test

import (
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetPortReturnsSpecificIfUsingEnvVar(t *testing.T) {
	_ = os.Setenv("PORT", "1234")
	assert.Equal(t, "1234", engine.GetPort())
	_ = os.Unsetenv("PORT")
}

func TestGetPortReturns8080(t *testing.T) {
	assert.Equal(t, "8080", engine.GetPort())
}

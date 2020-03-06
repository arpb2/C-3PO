package health_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/presentation/health"

	"github.com/arpb2/C-3PO/test/mock/http"
	"github.com/stretchr/testify/assert"
)

func TestHealthControllerBodyReturnsWith200IfItsOk(t *testing.T) {
	c, w := http.CreateTestContext()

	health.CreateGetHandler()(c)

	assert.Equal(t, 200, w.Code)
}

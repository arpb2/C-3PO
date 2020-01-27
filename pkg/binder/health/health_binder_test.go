package health_test

import (
	"testing"

	"github.com/arpb2/C-3PO/api/controller"
	healthbinder "github.com/arpb2/C-3PO/pkg/binder/health"
	"github.com/stretchr/testify/assert"
)

type MockControllerRegistrable struct {
	RegisteredControllers []controller.Controller
}

func (m *MockControllerRegistrable) Register(controller controller.Controller) {
	m.RegisteredControllers = append(m.RegisteredControllers, controller)
}

func bindControllers() *MockControllerRegistrable {
	binder := healthbinder.CreateBinder()
	registrable := &MockControllerRegistrable{}

	binder.BindControllers(registrable)

	return registrable
}

func lookupController(method, path string) *controller.Controller {
	registrable := bindControllers()
	var expectedController *controller.Controller

	for _, registeredController := range registrable.RegisteredControllers {
		if registeredController.Method == method && registeredController.Path == path {
			expectedController = &registeredController
		}
	}

	return expectedController
}

func TestCreateBinder_RegistersRoutes(t *testing.T) {
	assert.NotNil(t, lookupController("GET", "/ping"))
}

func TestCreateBinder_RegistersOnlyRoutes(t *testing.T) {
	registrable := bindControllers()

	assert.Equal(t, 1, len(registrable.RegisteredControllers))
}
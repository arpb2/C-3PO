package code_test

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/code"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockControllerRegistrable struct{
	RegisteredControllers []controller.Controller
}

func (m *MockControllerRegistrable) Register(controller controller.Controller) {
	m.RegisteredControllers = append(m.RegisteredControllers, controller)
}

func bindControllers() *MockControllerRegistrable {
	binder := code.CreateBinder(nil, nil)
	registrable := &MockControllerRegistrable{}

	binder.BindControllers(registrable)

	return registrable
}

func lookupController(path string) controller.Controller {
	registrable := bindControllers()
	var expectedController controller.Controller

	for _, registeredController := range registrable.RegisteredControllers {
		if registeredController.Path == path {
			expectedController = registeredController
		}
	}

	return expectedController
}

func TestCreateBinder_RegistersRoutes(t *testing.T) {
	assert.NotNil(t, lookupController("GET"))
	assert.NotNil(t, lookupController("POST"))
	assert.NotNil(t, lookupController("PUT"))
}

func TestCreateBinder_RegistersOnlyRoutes(t *testing.T) {
	registrable := bindControllers()

	assert.Equal(t, 3, len(registrable.RegisteredControllers))
}

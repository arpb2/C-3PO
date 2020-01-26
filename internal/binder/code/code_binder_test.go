package code_binder_test

import (
	controller2 "github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/internal/binder/code"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockControllerRegistrable struct {
	RegisteredControllers []controller2.Controller
}

func (m *MockControllerRegistrable) Register(controller controller2.Controller) {
	m.RegisteredControllers = append(m.RegisteredControllers, controller)
}

func bindControllers() *MockControllerRegistrable {
	binder := code_binder.CreateBinder(nil, nil, nil)
	registrable := &MockControllerRegistrable{}

	binder.BindControllers(registrable)

	return registrable
}

func lookupController(method, path string) *controller2.Controller {
	registrable := bindControllers()
	var expectedController *controller2.Controller

	for _, registeredController := range registrable.RegisteredControllers {
		if registeredController.Method == method && registeredController.Path == path {
			expectedController = &registeredController
		}
	}

	return expectedController
}

func TestCreateBinder_RegistersRoutes(t *testing.T) {
	assert.NotNil(t, lookupController("GET", "/users/:user_id/codes/:code_id"))
	assert.NotNil(t, lookupController("POST", "/users/:user_id/codes"))
	assert.NotNil(t, lookupController("PUT", "/users/:user_id/codes/:code_id"))
}

func TestCreateBinder_RegistersOnlyRoutes(t *testing.T) {
	registrable := bindControllers()

	assert.Equal(t, 3, len(registrable.RegisteredControllers))
}

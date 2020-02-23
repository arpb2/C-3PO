package binder_test

import (
	binder2 "github.com/arpb2/C-3PO/pkg/presentation/session/binder"
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/stretchr/testify/assert"
)

type MockControllerRegistrable struct {
	RegisteredControllers []controller.Controller
}

func (m *MockControllerRegistrable) Register(controller controller.Controller) {
	m.RegisteredControllers = append(m.RegisteredControllers, controller)
}

func bindControllers() *MockControllerRegistrable {
	binder := binder2.CreateBinder(nil, nil, nil)
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
	assert.NotNil(t, lookupController("POST", "/session"))
}

func TestCreateBinder_RegistersOnlyRoutes(t *testing.T) {
	registrable := bindControllers()

	assert.Equal(t, 1, len(registrable.RegisteredControllers))
}

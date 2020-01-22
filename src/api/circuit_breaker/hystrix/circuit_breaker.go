package hystrix

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/arpb2/C-3PO/src/api/circuit_breaker"
)

func CreateCircuitBreaker() circuit_breaker.CircuitBreaker {
	return &CircuitBreaker{}
}

type CircuitBreaker struct {}

func (cb CircuitBreaker) Go(command circuit_breaker.Command) chan error {
	return hystrix.Go(command.Name(), command.Run, command.Fallback)
}

func (cb CircuitBreaker) Do(command circuit_breaker.Command) error {
	return hystrix.Do(command.Name(), command.Run, command.Fallback)
}

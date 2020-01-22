package blocking

import (
	"github.com/arpb2/C-3PO/src/api/circuit_breaker"
)

func CreateCircuitBreaker() circuit_breaker.CircuitBreaker {
	return &CircuitBreaker{}
}

type CircuitBreaker struct {}

func (cb CircuitBreaker) Go(command circuit_breaker.Command) chan error {
	errChan := make(chan error, 1)
	err := command.Run()

	if err != nil {
		err = command.Fallback(err)
	}

	if err != nil {
		errChan <- err
	}
	return errChan
}

func (cb CircuitBreaker) Do(command circuit_breaker.Command) error {
	err := command.Run()

	if err != nil {
		return command.Fallback(err)
	}

	return err
}

package circuit_breaker

type Command interface {

	Name() string

	Run() error

	Fallback(err error) error

}

type CircuitBreaker interface {

	Go(command Command) chan error

	Do(command Command) error

}

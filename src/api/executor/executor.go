package executor

type Command interface {

	Name() string

	Run() error

	Fallback(err error) error

}

type Executor interface {

	Go(command Command) chan error

	Do(command Command) error

}

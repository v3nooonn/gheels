package notification

// ExecutorsOption the generation option of the Executors
type ExecutorsOption func(executors Executors)

// ExecutorFunc the generation of the Executor
type ExecutorFunc func() Executor

func WithExecutor(exeFunc ExecutorFunc) ExecutorsOption {
	return func(executors Executors) {
		executors.Add(exeFunc())
	}
}

// NewExecutors generation in functional option pattern
func NewExecutors(opts ...ExecutorsOption) Executors {
	notifier := new(Notifier)
	for _, option := range opts {
		option(notifier)
	}

	return notifier
}

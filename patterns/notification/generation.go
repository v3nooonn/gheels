package notification

// ExecutorsOption the generation option of the Executors
type ExecutorsOption func(executors Executors)

// NewExecutor the generation of the Executor
type NewExecutor func() Executor

func WithSendgrid(newSendgrid NewExecutor) ExecutorsOption {
	return func(executors Executors) {
		executors.Add(newSendgrid())
	}
}

func WithTwilio(newTwilio NewExecutor) ExecutorsOption {
	return func(executors Executors) {
		executors.Add(newTwilio())
	}
}

// NewExecutors generation in functional option pattern
func NewExecutors(options ...ExecutorsOption) Executors {
	notifier := new(Notifier)
	for _, option := range options {
		option(notifier)
	}

	return notifier
}

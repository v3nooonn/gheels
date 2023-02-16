package messager

type (
	// NotifierOption the generation option of the Executors
	NotifierOption func(notifier Executors)

	// ExeFuncBuilder builder of NewExecutor
	ExeFuncBuilder func(ro any) NewExecutor // any should be the `ReachOut` identified in proto

	// NewExecutor the generation function of the Executor
	NewExecutor func() Executor

	Mapper interface {
		Mapping() Executor
	}
)

func NewNotifier(opts ...NotifierOption) Executors {
	notifier := new(Notifier)
	for _, option := range opts {
		option(notifier)
	}

	return notifier
}

func WithExecutor(newExecutor NewExecutor) NotifierOption {
	return func(notifier Executors) {
		notifier.Add(newExecutor())
	}
}

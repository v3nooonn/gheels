// Package booster, all the lifecycle component within the project
// their initialization and termination are abstracted as a Booster.
// Which means, a component should provide a Booster or Wrapper
// instead of a specific function/method to init or terminate.
package booster

// Booster the abstraction interface
type Booster interface {
	Start()
	Stop()
}

// Booster wrapper
// Which wraps two functions into a Booster,
// that could be called as a single instance.
type (
	StartFunc func()
	StopFunc  func()

	Wrapper struct {
		StartFunc StartFunc
		StopFunc  StopFunc
	}
)

func NewWrapper(start StartFunc, stop StopFunc) Booster {
	return &Wrapper{StartFunc: start, StopFunc: stop}
}

func (b *Wrapper) Start() {
	if b.StartFunc != nil {
		b.StartFunc()
	}
}

func (b *Wrapper) Stop() {
	if b.StopFunc != nil {
		b.StopFunc()
	}
}

// StartsAll starts all the wrappers included
func StartsAll(wrappers []Booster) {
	for _, wrapper := range wrappers {
		wrapper.Start()
	}
}

// StopsAll stops all the wrappers included
func StopsAll(wrappers []Booster) {
	for _, wrapper := range wrappers {
		wrapper.Stop()
	}
}

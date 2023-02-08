package executors

import (
	"context"
	"fmt"
)

type Twilio struct {
	Target   string
	ParamMap map[string]string
}

func (t *Twilio) Execute(ctx context.Context) error {
	fmt.Printf("Executing Twilio: %s\n", t.Target)
	return nil
}

package email

import (
	"context"
	"fmt"
)

type Sendgrid struct {
	Target   string
	Template string
	ParamMap map[string]string
}

func (s *Sendgrid) Execute(ctx context.Context) error {
	fmt.Printf("Executing Sendgrid: %s\n", s.Target)
	return nil
}

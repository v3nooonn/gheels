package carrier

import (
	"github.com/v3nooonn/gheels/patterns/notification"
	"github.com/v3nooonn/gheels/patterns/notification/executors"
	"google.golang.org/protobuf/types/known/structpb"
)

type (
	Carrier struct {
		Subject   *Subject
		Audience  *Audience
		ReachOuts []*ReachOut
	}

	Subject struct{}

	Audience struct{}

	ReachOut struct {
		Via         Via
		Destination Destination
		ParamMap    *structpb.Struct
	}

	Via int8

	Destination struct{}
)

const (
	Unknown Via = iota
	Twilio
	Sendgrid
)

func (c *Carrier) Mapping() notification.Executor {
	// TODO: generate a specific Executor according to Carrier
	return &executors.Sendgrid{
		Target:   "",
		Template: "",
		ParamMap: nil,
	}
}

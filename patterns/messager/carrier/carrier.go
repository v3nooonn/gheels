package carrier

import (
	"github.com/v3nooonn/gheels/patterns/messager"
	"github.com/v3nooonn/gheels/patterns/messager/executors"

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

func (c *Carrier) Mapping() messager.Executor {
	// TODO: generate a specific Executor according to Carrier
	return &executors.Sendgrid{
		Target:   "",
		Template: "",
		ParamMap: nil,
	}
}

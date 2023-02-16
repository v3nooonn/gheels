package main

import (
	"context"

	"github.com/v3nooonn/gheels/patterns/messager"
	executors2 "github.com/v3nooonn/gheels/patterns/messager/executors"
)

func main() {
	pushNotification()
}

func pushNotification() {
	notifier := messager.NewNotifier(
		messager.WithExecutor(func() messager.Executor {
			return &executors2.Sendgrid{
				Target:   "email@gmail.com",
				Template: "template_id",
				ParamMap: map[string]string{
					"key": "value",
				},
			}
		}),
		messager.WithExecutor(func() messager.Executor {
			return &executors2.Twilio{
				Target: "+1111111111",
				ParamMap: map[string]string{
					"key": "value",
				},
			}
		}),
	)

	if err := notifier.Execute(context.TODO()); err != nil {
		panic(err)
	}
}

package main

import (
	"context"

	"github.com/v3nooonn/gheels/patterns/notification"
	"github.com/v3nooonn/gheels/patterns/notification/executors"
)

func main() {
	pushNotification()
}

func pushNotification() {
	notifier := notification.NewNotifier(
		notification.WithExecutor(func() notification.Executor {
			return &executors.Sendgrid{
				Target:   "email@gmail.com",
				Template: "template_id",
				ParamMap: map[string]string{
					"key": "value",
				},
			}
		}),
		notification.WithExecutor(func() notification.Executor {
			return &executors.Twilio{
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

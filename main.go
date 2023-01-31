package main

import (
	"context"

	"github.com/v3nooonn/gheels/patterns/notification"
	"github.com/v3nooonn/gheels/patterns/notification/executors/email"
	"github.com/v3nooonn/gheels/patterns/notification/executors/sms"
)

func main() {
	pushNotification()
}

func pushNotification() {
	notifier := notification.NewExecutors(
		notification.WithSendgrid(func() notification.Executor {
			return &email.Sendgrid{
				Target:   "email@gmail.com",
				Template: "template_id",
				ParamMap: map[string]string{
					"key": "value",
				},
			}
		}),
		notification.WithTwilio(func() notification.Executor {
			return &sms.Twilio{
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

// Package vendors ..
package vendors

import (
	"fmt"

	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

type NotificationData struct {
	ServiceID    string
	SubServiceID string
}

type NotificationMessage struct {
	Body  string
	Title string
	Data  NotificationData
}

func SendNotification(tokens []expo.ExponentPushToken, message NotificationMessage, data NotificationData) {

	// Create a new Expo SDK client
	client := expo.NewPushClient(nil)

	// Publish message
	response, err := client.Publish(
		&expo.PushMessage{
			To:       tokens,
			Body:     message.Body,
			Data:     map[string]string{"serviceID": data.ServiceID, "subServiceID": data.SubServiceID},
			Sound:    "default",
			Title:    message.Title,
			Priority: expo.DefaultPriority,
		},
	)
	// Check errors
	if err != nil {
		fmt.Println(err)
		return
	}
	// Validate responses
	if response.ValidateResponse() != nil {
		fmt.Println(response.PushMessage.To, "failed")
	}
}

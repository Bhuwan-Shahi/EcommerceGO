package notification

import (
	"ecommerceGO/config"
	"encoding/json"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type NotificationClient interface {
	SendSMS(phone, message string) error
}

type notificationClient struct {
	config config.Config
}

func (c notificationClient) SendSMS(phone, message string) error {

	accountSid := c.config.TwilioAccountSid
	authToken := c.config.TwilioAuthToken

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phone)                      //my phone customerphone
	params.SetFrom(c.config.TwilioFromPhone) //from gwilio
	params.SetBody(phone)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
	return nil
}

func NewNotificationCleint(config config.Config) NotificationClient {
	return &notificationClient{
		config: config,
	}
}

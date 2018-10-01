package sms

import (
	"encoding/json"
	"errors"
	"github.com/sfreiberg/gotwilio"
)

type TwilioSenderConfig struct {
	AccountSid string `json:"account_sid"`
	AuthToken  string `json:"auth_token"`
	FromPhone  string `json:"from_phone"`
}

type TwilioSender struct {
	config       *TwilioSenderConfig
	twilioClient *gotwilio.Twilio
}

// Todo: use NewTwilioClientWithCustomHttpClient to manager connection
func NewTwilioSender(config *TwilioSenderConfig) *TwilioSender {
	client := gotwilio.NewTwilioClientCustomHTTP(config.AccountSid, config.AuthToken, nil) // manager httpclient

	return &TwilioSender{
		config:       config,
		twilioClient: client,
	}
}

func (t *TwilioSender) SendSms(message SmsMessage) error {
	_, exception, err := t.twilioClient.SendSMS(t.config.FromPhone, message.Phone, message.Message, "", "")
	if err != nil {
		return err
	} else {
		if exception != nil {
			ex, _ := json.Marshal(exception)
			return errors.New(string(ex))
		}
	}
	return nil
}

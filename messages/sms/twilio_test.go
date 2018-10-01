package sms

import "testing"

func TestTwilioSender_SendSms(t *testing.T) {
	twilioSender := NewTwilioSender(&TwilioSenderConfig{
		AccountSid: "",
		AuthToken:  "",
		FromPhone:  "",
	})

	if err := twilioSender.SendSms(SmsMessage{
		Phone:   "",
		Message: "Test twilio.",
	}); err != nil {
		t.Error(err)
	} else {
		t.Log("Successful.")
	}
}

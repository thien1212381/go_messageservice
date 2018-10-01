package sms

type SmsMessage struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

type ISmsMessage interface {
	SendSms(message SmsMessage) error
}

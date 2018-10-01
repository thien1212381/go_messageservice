package email

import "testing"

func TestSMTPSender_SendEmail(t *testing.T) {
	smtpSender := NewSmtpSender(&SMTPSenderConfig{
		Sender: "",
		Pwd:    "",
		Server: "smtp.gmail.com",
		Port:   "587",
	})

	if err := smtpSender.SendEmail(EmailMessage{
		Destinations: []string{"youremail@gmail.com"},
		Subject:      "Test Send Email",
		Body:         "Test send email - Body",
		FromName:     "EmailTest",
	}); err != nil {
		t.Error(err)
	} else {
		t.Log("Successed")
	}
}

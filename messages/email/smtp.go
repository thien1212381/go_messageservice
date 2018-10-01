package email

import (
	"bytes"
	"net/smtp"
	"strings"
	"text/template"
)

type SMTPSenderConfig struct {
	Sender string `json:"sender"`
	Pwd    string `json:"pwd"`
	Server string `json:"server"`
	Port   string `json:"port"`
}

const DefaultEmailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}

Sincerely,

{{.FromName}}
`

type SMTPData struct {
	EmailMessage
	To   string
	From string
}

type SMTPSender struct {
	config       *SMTPSenderConfig
	textTemplate *template.Template
}

func NewSmtpSender(config *SMTPSenderConfig) *SMTPSender {
	t := template.New("PlainEmail")
	t, _ = t.Parse(DefaultEmailTemplate)

	return &SMTPSender{
		config:       config,
		textTemplate: t,
	}
}

// TODO: custom email template
func (s *SMTPSender) SendEmail(message EmailMessage) error {
	var msg bytes.Buffer

	if message.FromName == "" {
		message.FromName = s.config.Sender
	}

	data := &SMTPData{
		EmailMessage: message,
		To:           strings.Join(message.Destinations, ", "),
		From:         s.config.Sender,
	}

	err := s.textTemplate.Execute(&msg, data)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", s.config.Sender, s.config.Pwd, s.config.Server)
	serverAddress := s.config.Server + ":" + s.config.Port
	return smtp.SendMail(serverAddress, auth, s.config.Sender, message.Destinations, msg.Bytes())
}

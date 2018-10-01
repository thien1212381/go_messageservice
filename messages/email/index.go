package email

type EmailMessage struct {
	Destinations []string `json:"destinations"`
	Subject      string   `json:"subject"`
	Body         string   `json:"body"`
	FromAddress  string   `json:"from_address"`
	FromName     string   `json:"from_name"`
}

type IEmailMessage interface {
	SendEmail(config EmailMessage) error
}

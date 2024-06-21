package mail

import (
	"fmt"
	"github.com/resend/resend-go/v2"
)

type MailService struct {
	client *resend.Client
}

func NewMailService() (*MailService, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}
	client := resend.NewClient(cfg.ResendApiKey)
	return &MailService{client: client}, nil
}

func (m *MailService) SendMail() (string, error) {
	params := &resend.SendEmailRequest{
		From:    "Acme <onboarding@resend.dev>", // TODO: replace with arguments
		To:      []string{"kurocat2000@gmail.com"},
		Html:    "<strong>hello world</strong>",
		Subject: "Hello from Golang",
	}

	sent, err := m.client.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error()) // use logger
		return "", err
	}
	fmt.Printf("Email sent: %s\n", sent.Id)
	return sent.Id, nil
}

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

func (m *MailService) SendMail(fromName, to, html, subject string) (string, error) {
	params := &resend.SendEmailRequest{
		From:    fmt.Sprintf("%s <noreply@mail.cheesena.dev>", fromName),
		To:      []string{to},
		Html:    html,
		Subject: subject,
	}

	sent, err := m.client.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error()) // TODO: use logger
		return "", err
	}
	fmt.Printf("Email sent: %s\n", sent.Id)
	return sent.Id, nil
}

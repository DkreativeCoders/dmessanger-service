package mail

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mailgun/mailgun-go/v3"
	"os"
	"time"
)

func SendSimpleMessage(domain, apiKey string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		"Excited User <mailgun@YOUR_DOMAIN_NAME>",
		"Hello",
		"Testing some Mailgun awesomeness!",
		"YOU@YOUR_DOMAIN_NAME",
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}


type mailGunImplementation struct {
	domain string
	apiKey string

}


func NewMailGunImplementationNoArgs() IMail {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	domain := os.Getenv("MAIL_GUN_DOMAIN")
	apiKey := os.Getenv("MAIL_GUN_API_KEY")

	return mailGunImplementation{domain: domain, apiKey: apiKey}
}

func NewMailGunImplementation(domain string, apiKey string) IMail {
	return mailGunImplementation{domain: domain, apiKey: apiKey}
}

func (m mailGunImplementation) SendMail(subject, text string, to ...string) (string, error) {
	mg := mailgun.NewMailgun(m.domain, m.apiKey)

	message := mg.NewMessage(
		"dkreativecoders@gmail.com",
		subject,
		text,
		to...
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, message)
	return id, err
}

func (m mailGunImplementation) SendMailWithHtMlTemplate() (string, error) {
	panic("implement me")
}

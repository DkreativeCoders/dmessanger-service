package mail

import (
	"context"
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
	aoiKey string
}

func (m mailGunImplementation)  iniitalizeMailGun() (string, error) {

	panic("implement me")
}

func (m mailGunImplementation) SendMail() (string, error) {
	panic("implement me")
}

func (m mailGunImplementation) SendMailWithHtMlTemplate() (string, error) {
	panic("implement me")
}

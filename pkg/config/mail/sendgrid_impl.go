package mail

import (
	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

// sendGrid is send grid's implementation of the mail service
type sendGrid struct {
	apiKey string
}

func NewSendGridNoArgs() *sendGrid{
	err := godotenv.Load()
	if err != nil{
		panic(err)
	}
	apiKey := os.Getenv("SENDGRID_API_KEY")

	return &sendGrid{apiKey: apiKey}
}

func (s sendGrid) SendMail(subject, text string, to ...string) (string, error){
	// Email sender
	from := mail.NewEmail("Dkreative Coders", "dkreativecoders@gmail.com")

	// Email receiver
	recipient := mail.NewEmail("Customer", to[0])

	// Html text content
	htmlContent := text

	// New single mail
	message := mail.NewSingleEmail(from, subject, recipient, text, htmlContent)

	// Email sending client
	client := sendgrid.NewSendClient(s.apiKey)

	response, err := client.Send(message)
	if err != nil{
		return "", err
	}

	return response.Body, nil
}

func (s sendGrid) SendEMail(email EMailMessage) (string, error){
	// Email sender
	from := mail.NewEmail("Dkreative Coders", "dkreativecoders@gmail.com")

	// Email receiver
	recipient := mail.NewEmail("Customer", email.recipient)

	// Html text content
	htmlContent :=  email.text

	// New single mail
	message := mail.NewSingleEmail(from, email.subject, recipient, email.text, htmlContent)

	// Email sending client
	client := sendgrid.NewSendClient(s.apiKey)

	response, err := client.Send(message)
	if err != nil{
		return "", err
	}

	return response.Body, nil
}

func (s sendGrid) SendMailWithHtMlTemplate() (string, error){
	return "", nil
}
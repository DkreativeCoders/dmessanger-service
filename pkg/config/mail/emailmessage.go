package mail

type EMailMessage struct {
	subject string
	text string
	recipient string
	recipients [] string
}

func NewEMailMessage(subject string, text string, recipient string, recipients []string) *EMailMessage {
	return &EMailMessage{subject: subject, text: text, recipient: recipient, recipients: recipients}
}




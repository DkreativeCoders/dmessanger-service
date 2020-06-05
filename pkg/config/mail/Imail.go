package mail

type IMail interface {
	SendMail(subject, text string, to ...string) (string, error)
	SendEMail(email EMailMessage) (string, error)
	SendMailWithHtMlTemplate() (string, error)
}

package mail

type IMail interface {
	SendMail() (string, error)
	SendMailWithHtMlTemplate() (string, error)
}
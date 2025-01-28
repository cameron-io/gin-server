package interfaces

type MailService interface {
	MailLoginToken(token string)
}

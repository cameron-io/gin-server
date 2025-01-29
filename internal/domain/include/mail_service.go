package include

type MailService interface {
	MailLoginToken(token string)
}

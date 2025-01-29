package services

import (
	"os"

	"cameron.io/gin-server/internal/domain/include"
	"cameron.io/gin-server/pkg/mail"
)

type MailService struct {
}

func NewMailService() include.MailService {
	return &MailService{}
}

func (s *MailService) MailLoginToken(token string) {
	mail.Send(os.Getenv("EMAIL_SENDER"), "Login", `
		<html>
			<body>
				<a href="http://localhost:5000/api/accounts/login?token=`+token+`">Login Link</a>
			</body>
		</html>
	`)
}

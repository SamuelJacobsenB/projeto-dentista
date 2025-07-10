package utils

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/config"
)

const from = config.EMAIL
const password = config.EMAIL_APP_PASSWORD

var smtpHost string = fmt.Sprintf("smtp.%s.com", config.MAIL_PROVIDER)

const smtpPort = config.SMTP_PORT

func SendEmail(to string, subject string, body string) error {
	msg := []byte("Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg); err != nil {
		return fmt.Errorf("erro ao enviar email: %s", err.Error())
	}

	return nil
}

func GenerateBodyText(startTime, endTime time.Time) string {
	date := startTime.Format("01/01/2000")

	return fmt.Sprintf(
		"Olá! Este é um lembrete da sua consulta marcada para o dia %s, das %s às %s.\n\nPor favor, chegue com 10 minutos de antecedência. Caso precise remarcar, entre em contato conosco.\n\nAté breve!",
		date, startTime.Format("01:00"), endTime.Format("01:00"),
	)
}

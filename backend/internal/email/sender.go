package email

import (
	"fmt"
	"net/smtp"
)

func SendRequest(email Email, smtpHost string, smtpPort string) error {
	address := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	message := []byte(
		"From: " + email.SenderName + " <" + email.SenderEmail + ">\r\n" +
			"To: " + email.RecipientEmail + "\r\n" +
			"Subject: " + email.Subject + "\r\n" +
			"\r\n" +
			email.Content + "\r\n",
	)

	err := smtp.SendMail(address, nil, email.SenderEmail, []string{email.RecipientEmail}, message)

	if err != nil {
	return fmt.Errorf("Failed to send anonymous mail : %w", err)
}
	return nil
}
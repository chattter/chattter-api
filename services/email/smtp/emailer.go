package smtp

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/chattter/chattter-api/services/email"
)

type Emailer struct {
	Host     string
	Port     uint16
	Username string
	Password string
}

func (e *Emailer) Email(options *email.SendOptions) error {

	// Normalize the port number for the emailer
	port := e.Port
	if port == 0 {
		port = 25
	}

	// Format the full host and port address
	addr := fmt.Sprintf("%s:%d", e.Host, e.Port)

	// Create the SMTP authentication data
	auth := smtp.PlainAuth("", e.Username, e.Password, e.Host)

	// Get the strings for all recipience
	var tos []string
	for _, to := range options.To {
		tos = append(tos, to.String())
	}

	// Create the raw message
	msgParts := []string{
		fmt.Sprintf("To: %s", strings.Join(tos, ",")),
		fmt.Sprintf("Subject: %s", options.Subject),
		"",
		options.Body,
	}
	msg := strings.Join(msgParts, "\r\n")

	// Send the email
	return smtp.SendMail(
		addr,
		auth,
		options.From.String(),
		tos,
		[]byte(msg),
	)

}

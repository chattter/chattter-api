package services

import (
	"os"
	"strconv"

	"github.com/chattter/chattter-api/services/email"
	"github.com/chattter/chattter-api/services/email/smtp"
)

// EmailService provides email sending functionality
type EmailService struct{}

// GetEmail gets an emailer instance that is configured for the platform
func (s *EmailService) GetEmailer() (email.Emailer, error) {

	// Get the email provider
	provider := os.Getenv("EMAIL_PROVIDER")

	// If the email provider is SMTP
	if provider == "smtp" {

		// Parse the port number to a uint16
		var port uint16
		portStr := os.Getenv("SMTP_PORT")
		if len(portStr) > 0 {
			portNum, err := strconv.ParseInt(portStr, 10, 16)
			if err != nil {
				return nil, err
			}
			port = uint16(portNum)
		}

		// Return the SMTP emailer
		return &smtp.Emailer{
			Host:     os.Getenv("SMTP_HOST"),
			Port:     port,
			Username: os.Getenv("SMTP_USERNAME"),
			Password: os.Getenv("SMTP_PASSWORd"),
		}, nil

	}

	// If we get to this point, no provider option is configured, so return nil
	return nil, nil

}

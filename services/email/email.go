package email

import "errors"

// EmailService provides email sending functionality
type EmailService struct{}

// GetEmail gets an emailer instance that is configured for the platform
func (s *EmailService) GetEmailer() (Emailer, error) {
	return nil, errors.New("GetEmailer not implemented yet")
}

package ses

import (
	"errors"

	"github.com/chattter/chattter-api/services/email"
)

type Emailer struct{}

func (e *Emailer) Email(options *email.SendOptions) error {
	return errors.New("ses Email(...) not implemented")
}

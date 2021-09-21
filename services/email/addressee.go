package email

import "fmt"

type Addressee struct {
	Name  string
	Email string
}

func (a *Addressee) String() string {
	if len(a.Name) == 0 {
		return a.Email
	}
	return fmt.Sprintf("%s <%s>", a.Name, a.Email)
}

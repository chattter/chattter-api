package email

type SendOptions struct {
	To      []Addressee
	From    Addressee
	Subject string
	Body    string
}

type Emailer interface {
	Email(options *SendOptions) error
}

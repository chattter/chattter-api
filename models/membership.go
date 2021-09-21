package models

type Membership struct {
	AccountID uint64
	Account   *Account
	ChannelID uint64
	Channel   *Channel
}

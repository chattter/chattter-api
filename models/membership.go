package models

import (
	"database/sql"
	"time"
)

type Membership struct {
	ID          uint64 `gorm:"primaryKey"`
	AccountID   uint64
	Account     *Account
	ChannelID   uint64
	Channel     *Channel
	CreatedDate time.Time
	DeletedDate sql.NullTime
}

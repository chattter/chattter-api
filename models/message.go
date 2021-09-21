package models

import (
	"database/sql"
	"time"
)

type Message struct {
	ID          uint64 `gorm:"primaryKey"`
	AccountID   uint64
	Account     *Account
	ChannelID   uint64
	Channel     *Channel
	Message     string
	CreatedDate time.Time
	DeletedDate sql.NullTime
}

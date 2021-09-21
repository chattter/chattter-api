package models

import (
	"database/sql"
	"time"
)

type Channel struct {
	ID          uint64 `gorm:"primaryKey"`
	Name        string
	CreatedDate time.Time
	DeletedDate sql.NullTime
}

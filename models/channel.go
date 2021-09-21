package models

import (
	"database/sql"
	"image"
	"time"
)

type Channel struct {
	ID          uint64 `gorm:"primaryKey"`
	Name        string
	Logo        image.Image
	CreatedDate time.Time
	DeletedDate sql.NullTime
}

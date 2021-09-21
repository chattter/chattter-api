package models

import "image"

type Channel struct {
	ID          uint64 `gorm:"primaryKey"`
	ChannelName string
	Logo        image.Image
}

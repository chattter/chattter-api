package services

import (
	"errors"

	"gorm.io/gorm"
)

type ChatService struct {
	DB *gorm.DB
}

func (s *ChatService) SendMessage() error {
	return errors.New("SendMessage not implemented")
}

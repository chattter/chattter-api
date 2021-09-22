package services

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/chattter/chattter-api/models"
	"gorm.io/gorm"
)

type ChatService struct {
	DB *gorm.DB
}

func (s *ChatService) SendMessage(text string, channelID uint64, senderID uint64) (*models.Message, error) {
	message := models.Message{
		AccountID:   senderID,
		ChannelID:   channelID,
		Message:     text,
		CreatedDate: time.Now(),
	}
	if err := s.DB.Create(&message).Error; err != nil {
		return nil, err
	}
	if err := validateMessage(text); err != nil {
		return nil, err
	}
	return &message, nil

}

func (s *ChatService) DeleteMessage(id uint64, accountID uint64) error {
	return s.DB.
		Where("id = ?", id).
		Where("account_id = ?", accountID).
		Update(
			"deleted_date",
			sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		).
		Error
}

func (s *ChatService) EditMessage(id uint64, accountID uint64, message string) error {
	if err := validateMessage(message); err != nil {
		return err
	}
	return s.DB.
		Where("id = ?", id).
		Where("account_id = ?", accountID).
		Update(
			"message",
			message,
		).
		Error
}

func validateMessage(message string) error {
	message = strings.TrimSpace(message)
	if len(message) == 0 {
		return errors.New("message is too short")
	}
	return nil
}

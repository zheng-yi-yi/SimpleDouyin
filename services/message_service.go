package services

import (
	"time"

	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

type MessageService struct {
}

func (MessageService *MessageService) AddMessage(fromUserID, toUserID uint, content string) error {
	curMessage := models.Message{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Content:    content,
		CreateTime: time.Now(),
	}
	err := config.DB.Create(&curMessage).Error
	if err != nil {
		return err
	}
	return nil
}

func (MessageService *MessageService) MessageList(fromUserID, toUserID uint) ([]models.Message, error) {
	message_list := make([]models.Message, 0)
	config.DB.Where("from_user_id = ? AND to_user_id = ?", fromUserID, toUserID).Find(&message_list)
	return message_list, nil
}

package services

import (
	"time"

	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

type MessageService struct {
}

func (ms *MessageService) AddMessage(fromUserID uint, content string, toUserID uint) error {
	// 创建消息结构体
	message := models.Message{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Content:    content,
		CreateTime: time.Now(),
	}
	// 将消息插入到数据库
	if err := config.Database.Create(&message).Error; err != nil {
		return err
	}

	return nil
}

func (MessageService *MessageService) GetMessageListWithTime(fromUserID, toUserID uint, preMsgTime time.Time) ([]models.Message, error) {
	var messages []models.Message

	// 查询消息记录，基于 pre_msg_time 进行查询
	err := config.Database.Where("((from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)) AND create_time > ?",
		fromUserID, toUserID, toUserID, fromUserID, preMsgTime).
		Order("create_time").
		Find(&messages).Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}

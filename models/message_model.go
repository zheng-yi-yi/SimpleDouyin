package models

import (
	"time"
)

// @Description: 消息模型
type Message struct {
	ID         uint      `json:"id" gorm:"primaryKey comment:消息id"`
	FromUserID uint      `json:"from_user_id" gorm:"not null comment:消息发送者id"`
	ToUserID   uint      `json:"to_user_id" gorm:"not null comment:消息接收者id"`
	Content    string    `json:"content" gorm:"not null comment:消息内容"`
	CreateTime time.Time `json:"create_time" gorm:"not null comment:消息发送时间"`
}

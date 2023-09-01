package models

import (
	"time"
)

// Message, 消息表
type Message struct {
	ID         uint      `gorm:"primaryKey comment:消息id"`
	FromUserID uint      `gorm:"not null   comment:消息发送者id; type:INT"`
	ToUserID   uint      `gorm:"not null   comment:消息接收者id; type:INT"`
	Content    string    `gorm:"not null   comment:消息内容;     type:TEXT"`
	CreateTime time.Time `gorm:"not null   comment:消息发送时间; type:DATETIME"`
}

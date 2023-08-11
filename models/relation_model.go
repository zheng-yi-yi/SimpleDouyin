package models

import "gorm.io/gorm"

type Relation struct {
	gorm.Model
	FromUserId uint `json:"from_user_id" gorm:"not null; comment: 关注人ID"`
	ToUserId   uint `json:"to_user_id" gorm:"not null; comment: 被关注人ID"`
	IsMutual   uint `json:"is_mutual" gorm:"type: tinyint(1); not null; default: 0; comment: 是否互相关注"`
}

package models

type Relation struct {
	ID         uint `gorm:"primarykey"`
	FromUserId uint `gorm:"not null; comment: 用户id"`
	ToUserId   uint `gorm:"not null; comment: 关注的用户"`
	Cancel     uint `gorm:"type: tinyint(1); not null; default: 0; comment: 默认关注为0，取消关注为1"`
	IsMutual   uint `gorm:"type: tinyint(1); not null; default: 0; comment: 是否互相关注"`
}

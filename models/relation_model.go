package models

// Relation, 关注关系表
type Relation struct {
	ID         uint `gorm:"primarykey"`
	FromUserId uint `gorm:"not null; comment: 用户id;                  type:INT"`
	ToUserId   uint `gorm:"not null; comment: 关注的用户;               type:INT"`
	Cancel     uint `gorm:"not null; comment: 默认关注为0，取消关注为1;  type:INT"`
}

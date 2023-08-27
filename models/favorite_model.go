package models

// Favorite, 点赞表
type Favorite struct {
	ID      uint `gorm:"primarykey"`
	UserId  uint `gorm:"not null; comment:用户ID"`
	VideoId uint `gorm:"not null; comment:视频ID"`
	Status  uint `gorm:"not null; comment:0取消点赞,1点赞(默认1)"`
	// 定义外键关系
	User  User  `gorm:"foreignKey:UserId;  references:ID;"`
	Video Video `gorm:"foreignKey:VideoId; references:ID;"`
}

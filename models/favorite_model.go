package models

// Favorite, 点赞表
type Favorite struct {
	ID      uint `gorm:"primarykey"`
	UserId  uint `gorm:"not null; comment:用户ID;  type:INT"`
	VideoId uint `gorm:"not null; comment:视频ID;  type:INT"`
	// 定义外键关系
	User  User  `gorm:"foreignKey:UserId;  references:ID; comment:点赞用户的信息"`
	Video Video `gorm:"foreignKey:VideoId; references:ID; comment:点赞视频的信息"`
}

package models

// User, 用户表
type User struct {
	ID              uint   `gorm:"primarykey"`
	UserName        string `gorm:"not null; comment:用户名"`
	PassWord        string `gorm:"not null; comment:密码"`
	Status          int    `gorm:"not null; comment:用户状态 0禁用 1启用 默认启用;"`
	FollowCount     int    `gorm:"not null; comment:关注总数"`
	FollowerCount   int    `gorm:"not null; comment:粉丝总数"`
	FavoriteCount   int64  `gorm:"not null; comment:喜欢数"`
	Avatar          string `gorm:"not null; comment:用户头像"`
	BackgroundImage string `gorm:"not null; comment:用户个人页顶部大图"`
	Signature       string `gorm:"not null; comment:个人简介"`
	TotalFavorited  string `gorm:"not null; comment:获赞数量"`
	WorkCount       int64  `gorm:"not null; comment:作品数"`
	// 定义外键关系
	Video   []Video   `gorm:"foreignKey:UserId; references:ID; comment:视频作者信息"`
	Comment []Comment `gorm:"foreignKey:UserId; references:ID; comment:评论信息"`
}

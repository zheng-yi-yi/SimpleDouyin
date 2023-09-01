package models

// User, 用户表
type User struct {
	ID              uint   `gorm:"primarykey"`
	UserName        string `gorm:"not null; comment:用户名;   type:VARCHAR(255)"`
	PassWord        string `gorm:"not null; comment:密码;     type:VARCHAR(255)"`
	FollowCount     int    `gorm:"not null; comment:关注总数; type:INT"`
	FollowerCount   int    `gorm:"not null; comment:粉丝总数; type:INT"`
	FavoriteCount   int64  `gorm:"not null; comment:喜欢数;   type:BIGINT"`
	TotalFavorited  string `gorm:"not null; comment:获赞数量; type:VARCHAR(255)"`
	WorkCount       int64  `gorm:"not null; comment:作品数;   type:BIGINT"`
	Avatar          string `gorm:"not null; comment:用户头像; type:VARCHAR(255)"`
	BackgroundImage string `gorm:"not null; comment:顶部图;   type:VARCHAR(255)"`
	Signature       string `gorm:"not null; comment:个人简介; type:TEXT"`
	// 定义外键关系
	Video   []Video   `gorm:"foreignKey:UserId; references:ID; comment:视频信息"`
	Comment []Comment `gorm:"foreignKey:UserId; references:ID; comment:评论信息"`
}

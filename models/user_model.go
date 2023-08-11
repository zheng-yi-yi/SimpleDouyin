package models

import "gorm.io/gorm"

//  @Description: 用户模型
type User struct {
	gorm.Model
	UserName      string    `json:"username,omitempty" gorm:"type:varchar(32); column:username; not null; unique; comment:用户名"`
	PassWord      string    `json:"-"                  gorm:"type:varchar(32); column:password; not null; comment:密码"`
	Status        int       `json:"status,omitempty"   gorm:"type:tinyint(1); default:1; not null; comment:用户状态 0禁用 1启用 默认启用;"`
	FollowCount   int       `json:"follow_count,omitempty" gorm:"default:0; not null; comment: 用户关注数"`
	FollowerCount int       `json:"follower_count,omitempty" gorm:"default:0; not null; comment: 用户粉丝数"`
	IsFollow      bool      `json:"is_follow,omitempty" gorm:"-"`
	Video         []Video   `gorm:"foreignKey:UserId"`
	Comment       []Comment `gorm:"foreignKey:UserId"`
}

package models

import "gorm.io/gorm"

//  @Description: 用户模型
type User struct {
	ID              uint      `json:"id"                         gorm:"primarykey"`
	UserName        string    `json:"name,omitempty"             gorm:"not null; unique; comment:用户名"`
	PassWord        string    `json:"-"                          gorm:"not null; comment:密码"`
	Status          int       `json:"status,omitempty"           gorm:"not null; default:1; comment:用户状态 0禁用 1启用 默认启用;"`
	FollowCount     int       `json:"follow_count,omitempty"     gorm:"not null; default:0; comment: 关注总数"`
	FollowerCount   int       `json:"follower_count,omitempty"   gorm:"not null; default:0; comment: 粉丝总数"`
	FavoriteCount   int64     `json:"favorite_count,omitempty"   gorm:"not null; default:0; comment:喜欢数"`
	Avatar          string    `json:"avatar,omitempty"           gorm:"not null; comment:用户头像"`
	BackgroundImage string    `json:"background_image,omitempty" gorm:"not null; comment:用户个人页顶部大图"`
	Signature       string    `json:"signature,omitempty"        gorm:"not null; comment:个人简介"`
	TotalFavorited  string    `json:"total_favorited,omitempty"  gorm:"not null; comment:获赞数量"`
	WorkCount       int64     `json:"work_count,omitempty"       gorm:"not null; comment:作品数"`
	Video           []Video   `gorm:"foreignKey:UserId"`
	Comment         []Comment `gorm:"foreignKey:UserId"`
}

// IncrementWorkCount 增加用户的作品数。
func IncrementWorkCount(db *gorm.DB, userID uint) error {
	return db.Model(&User{}).Where("id = ?", userID).Update("work_count", gorm.Expr("work_count + ?", 1)).Error
}

// IncrementUserLikeCount 增加用户的点赞数。
func IncrementUserLikeCount(db *gorm.DB, userID uint) error {
	return db.Model(&User{}).Where("id = ?", userID).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
}

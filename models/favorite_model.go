package models

type Favorite struct {
	ID      uint  `gorm:"primarykey"`
	UserId  uint  `json:"user_id,omitempty"  gorm:"not null; comment:用户ID"`
	VideoId uint  `json:"video_id,omitempty" gorm:"not null; comment:视频ID"`
	Status  uint  `json:"status,omitempty"   gorm:"default: 1; type:tinyint(2); not null; comment: 是否点赞 0取消点赞 1点赞 默认1"`
	User    User  `json:"user"               gorm:"foreignKey:UserId; references:ID;"`
	Video   Video `json:"video"              gorm:"foreignKey:VideoId; references:ID;"`
}

package config

import (
	"github.com/zheng-yi-yi/simpledouyin/models"
	"gorm.io/gorm"
)

// CreateTable ：自动创建（或更新）数据库表格，并初始化表格数据
func CreateTable(db *gorm.DB) error {
	err := db.AutoMigrate(
		models.User{},
		models.Video{},
		models.Favorite{},
		models.Comment{},
		models.Relation{},
		models.Message{},
	)
	if err != nil {
		return err
	}
	// 初始化用户表：
	var count_User int64
	db.Model(&models.User{}).Count(&count_User)
	if count_User == 0 {
		user1 := models.User{
			UserName:      "User1",
			PassWord:      "123456",
			Status:        1,
			FollowCount:   0,
			FollowerCount: 1,
		}
		db.Create(&user1)
		user2 := models.User{
			UserName:      "User2",
			PassWord:      "123456",
			Status:        1,
			FollowCount:   1,
			FollowerCount: 0,
		}
		db.Create(&user2)
	}
	// 初始化视频表：
	var count_Video int64
	db.Model(&models.Video{}).Count(&count_Video)
	if count_Video == 0 {
		videos := []*models.Video{
			{UserId: 1, PlayUrl: "videos/1.mp4", CoverUrl: "images/1.jpg", FavoriteCount: 1, Description: "邀你观看2023稀土开发者大会!11场论坛，60位技术大咖，涵盖AIGC与大模型、大前端、音视频等前沿技术资讯"},
			{UserId: 1, PlayUrl: "videos/2.mp4", CoverUrl: "images/2.jpg", FavoriteCount: 1, Description: "那些仅凭半句就封神的诗句"},
			{UserId: 1, PlayUrl: "videos/3.mp4", CoverUrl: "images/3.jpg", FavoriteCount: 1, Description: "登昆仑兮食玉英，与天地兮同寿与日月兮齐光。#汉服之美在华夏"},
			{UserId: 1, PlayUrl: "videos/4.mp4", CoverUrl: "images/4.jpg", FavoriteCount: 1, Description: "落日沉溺于橘色的海，晚风沦陷于赤城的爱"},
			{UserId: 1, PlayUrl: "videos/5.mp4", CoverUrl: "images/5.jpg", FavoriteCount: 1, Description: "五档 启动！ #太阳神尼卡登场"},
			{UserId: 1, PlayUrl: "videos/6.mp4", CoverUrl: "images/6.jpg", FavoriteCount: 1, Description: "好像做了一场短暂的山水梦"},
		}
		db.Create(&videos)
	}
	// 初始化点赞表：
	var count_Favorite int64
	db.Model(&models.Favorite{}).Count(&count_Favorite)
	if count_Favorite == 0 {
		favorite := []*models.Favorite{
			{UserId: 1, VideoId: 1, Status: 1},
			{UserId: 1, VideoId: 2, Status: 1},
			{UserId: 1, VideoId: 3, Status: 1},
			{UserId: 1, VideoId: 4, Status: 1},
			{UserId: 1, VideoId: 5, Status: 1},
		}
		db.Create(&favorite)
	}
	// 初始化评论表：
	var count_Comment int64
	db.Model(&models.Comment{}).Count(&count_Comment)
	if count_Comment == 0 {
		comment := []*models.Comment{
			{VideoId: 1, Content: "来学习了", ToUserId: 1, UserId: 2},
			{VideoId: 2, Content: "为天地立心，为生民立命，为往圣继绝学，为万世开太平", ToUserId: 1, UserId: 2},
			{VideoId: 3, Content: "变装那一瞬间好高级", ToUserId: 1, UserId: 2},
			{VideoId: 4, Content: "我对着日落许愿，希望你永远快乐", ToUserId: 1, UserId: 2},
			{VideoId: 5, Content: "年少的梦终将绽放于盛夏，解放之鼓必将响彻整个夏天", ToUserId: 1, UserId: 2},
			{VideoId: 5, Content: "说走就走的旅行...", ToUserId: 1, UserId: 2},
		}
		db.Create(&comment)
	}
	// 初始化关注关系表：
	var count_Relation int64
	db.Model(&models.Relation{}).Count(&count_Relation)
	if count_Relation == 0 {
		relation := []*models.Relation{
			{FromUserId: 2, ToUserId: 1, IsMutual: 0},
		}
		db.Create(&relation)
	}
	return nil
}

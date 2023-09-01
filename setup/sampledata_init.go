package setup

import (
	"time"

	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
	"github.com/zheng-yi-yi/simpledouyin/utils"
	"gorm.io/gorm"
)

func getpwd() []byte {
	encryptedPassword, _ := utils.EncryptPassword("123456")
	return encryptedPassword
}

// ========== 用户表-样例数据 ==========

var initialUsers = []models.User{
	{
		UserName:        "User1",                           // 用户名
		PassWord:        string(getpwd()),                  // 密码
		FollowCount:     3,                                 // 关注总数
		FollowerCount:   5,                                 // 粉丝总数
		FavoriteCount:   6,                                 // 喜欢数
		Avatar:          config.USER1_AVATAR_URL,           // 用户头像
		BackgroundImage: config.USER1_BACKGROUND_IMAGE_URL, // 用户个人页顶部大图
		Signature:       config.USER1_PROFILE_DESCRIPTION,  // 个人简介
		TotalFavorited:  "8",                               // 获赞数量
		WorkCount:       2,                                 // 作品数
	},
	{
		UserName:        "User2",                           // 用户名
		PassWord:        string(getpwd()),                  // 密码
		FollowCount:     3,                                 // 关注总数
		FollowerCount:   3,                                 // 粉丝总数
		FavoriteCount:   7,                                 // 喜欢数
		Avatar:          config.USER2_AVATAR_URL,           // 用户头像
		BackgroundImage: config.USER2_BACKGROUND_IMAGE_URL, // 用户个人页顶部大图
		Signature:       config.USER2_PROFILE_DESCRIPTION,  // 个人简介
		TotalFavorited:  "4",                               // 获赞数量
		WorkCount:       1,                                 // 作品数
	},
	{
		UserName:        "User3",                           // 用户名
		PassWord:        string(getpwd()),                  // 密码
		FollowCount:     4,                                 // 关注总数
		FollowerCount:   3,                                 // 粉丝总数
		FavoriteCount:   7,                                 // 喜欢数
		Avatar:          config.USER3_AVATAR_URL,           // 用户头像
		BackgroundImage: config.USER3_BACKGROUND_IMAGE_URL, // 用户个人页顶部大图
		Signature:       config.USER3_PROFILE_DESCRIPTION,  // 个人简介
		TotalFavorited:  "6",                               // 获赞数量
		WorkCount:       2,                                 // 作品数
	},
	{
		UserName:        "User4",                           // 用户名
		PassWord:        string(getpwd()),                  // 密码
		FollowCount:     4,                                 // 关注总数
		FollowerCount:   4,                                 // 粉丝总数
		FavoriteCount:   6,                                 // 喜欢数
		Avatar:          config.USER4_AVATAR_URL,           // 用户头像
		BackgroundImage: config.USER4_BACKGROUND_IMAGE_URL, // 用户个人页顶部大图
		Signature:       config.USER4_PROFILE_DESCRIPTION,  // 个人简介
		TotalFavorited:  "7",                               // 获赞数量
		WorkCount:       2,                                 // 作品数
	},
	{
		UserName:        "User5",                           // 用户名
		PassWord:        string(getpwd()),                  // 密码
		FollowCount:     3,                                 // 关注总数
		FollowerCount:   3,                                 // 粉丝总数
		FavoriteCount:   6,                                 // 喜欢数
		Avatar:          config.USER5_AVATAR_URL,           // 用户头像
		BackgroundImage: config.USER5_BACKGROUND_IMAGE_URL, // 用户个人页顶部大图
		Signature:       config.USER5_PROFILE_DESCRIPTION,  // 个人简介
		TotalFavorited:  "11",                              // 获赞数量
		WorkCount:       2,                                 // 作品数
	},
	{
		UserName:        "User6",                           // 用户名
		PassWord:        string(getpwd()),                  // 密码
		FollowCount:     4,                                 // 关注总数
		FollowerCount:   3,                                 // 粉丝总数
		FavoriteCount:   7,                                 // 喜欢数
		Avatar:          config.USER6_AVATAR_URL,           // 用户头像
		BackgroundImage: config.USER6_BACKGROUND_IMAGE_URL, // 用户个人页顶部大图
		Signature:       config.USER6_PROFILE_DESCRIPTION,  // 个人简介
		TotalFavorited:  "4",                               // 获赞数量
		WorkCount:       1,                                 // 作品数
	},
}

func initUsers(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count == 0 {
		for _, user := range initialUsers {
			db.Create(&user)
		}
	}
}

// ========== 视频表-样例数据 ==========

var initialVideos = []models.Video{
	{
		UserId:        1,                // 用户id
		PlayUrl:       "videos/1_1.mp4", // 视频地址
		CoverUrl:      "images/1_1.png", // 封面图地址
		FavoriteCount: 5,                // 点赞数量
		CommentCount:  3,                // 评论数量
		Description:   "回忆这把刀",          // 视频标题
		CreatedAt:     time.Now(),
	},
	{
		UserId:        5,                // 用户id
		PlayUrl:       "videos/5_1.mp4", // 视频地址
		CoverUrl:      "images/5_1.png", // 封面图地址
		FavoriteCount: 5,                // 点赞数量
		CommentCount:  3,                // 评论数量
		Description:   "而我独缺你一生的了解 #戏曲", // 视频标题
		CreatedAt:     time.Now(),
	},
	{
		UserId:        3,                      // 用户id
		PlayUrl:       "videos/3_1.mp4",       // 视频地址
		CoverUrl:      "images/3_1.png",       // 封面图地址
		FavoriteCount: 4,                      // 点赞数量
		CommentCount:  2,                      // 评论数量
		Description:   "哪有什么突然好想你 明明就心里一直有你.", // 视频标题
		CreatedAt:     time.Now(),
	},
	{
		UserId:        6,                // 用户id
		PlayUrl:       "videos/6_1.mp4", // 视频地址
		CoverUrl:      "images/6_1.png", // 封面图地址
		FavoriteCount: 4,                // 点赞数量
		CommentCount:  3,                // 评论数量
		Description:   "好像做了一场短暂的山水梦",   // 视频标题
		CreatedAt:     time.Now(),
	},
	{
		UserId:        1,                               // 用户id
		PlayUrl:       "videos/1_2.mp4",                // 视频地址
		CoverUrl:      "images/1_2.png",                // 封面图地址
		FavoriteCount: 3,                               // 点赞数量
		CommentCount:  2,                               // 评论数量
		Description:   "登昆仑兮食玉英，与天地兮同寿与日月兮齐光。#汉服之美在华夏", // 视频标题
		CreatedAt:     time.Now(),
	},
	{
		UserId:        4,                // 用户id
		PlayUrl:       "videos/4_1.mp4", // 视频地址
		CoverUrl:      "images/4_1.png", // 封面图地址
		FavoriteCount: 4,                // 点赞数量
		CommentCount:  1,                // 评论数量
		Description:   "世上有无条件的爱吗？",     // 视频标题
		CreatedAt:     time.Now(),
	},
	{
		UserId:        4,                 // 用户id
		PlayUrl:       "videos/4_2.mp4",  // 视频地址
		CoverUrl:      "images/4_2.png",  // 封面图地址
		FavoriteCount: 3,                 // 点赞数量
		CommentCount:  3,                 // 评论数量
		Description:   "五档 启动！ #太阳神尼卡登场", // 视频标题
		CreatedAt:     time.Now(),
	},
	{
		UserId:        2,                // 用户id
		PlayUrl:       "videos/2_1.mp4", // 视频地址
		CoverUrl:      "images/2_1.png", // 封面图地址
		FavoriteCount: 3,                // 点赞数量
		CommentCount:  2,                // 评论数量
		Description:   "那些仅凭半句就封神的诗句",   // 视频标题
		CreatedAt:     time.Now(),
	},
	{
		UserId:        3,                     // 用户id
		PlayUrl:       "videos/3_2.mp4",      // 视频地址
		CoverUrl:      "images/3_2.png",      // 封面图地址
		FavoriteCount: 2,                     // 点赞数量
		CommentCount:  1,                     // 评论数量
		Description:   "落日沉溺于橘色的海，晚风沦陷于赤城的爱", // 视频标题
		CreatedAt:     time.Now(),
	},
	{
		UserId:        5,                            // 用户id
		PlayUrl:       "videos/5_2.mp4",             // 视频地址
		CoverUrl:      "images/5_2.png",             // 封面图地址
		FavoriteCount: 6,                            // 点赞数量
		CommentCount:  4,                            // 评论数量
		Description:   "文有太极安天下，武有八极定乾坤 #太极拳 #国风少年", // 视频标题
		CreatedAt:     time.Now(),
	},
}

func initVideos(db *gorm.DB) {
	var count int64
	db.Model(&models.Video{}).Count(&count)
	if count == 0 {
		for _, video := range initialVideos {
			db.Create(&video)
		}
	}
}

// ========== 点赞表-样例数据 ==========

var initialFavorites = []models.Favorite{
	{UserId: 2, VideoId: 1},
	{UserId: 3, VideoId: 1},
	{UserId: 4, VideoId: 1},
	{UserId: 5, VideoId: 1},
	{UserId: 6, VideoId: 1},
	{UserId: 1, VideoId: 2},
	{UserId: 5, VideoId: 2},
	{UserId: 4, VideoId: 2},
	{UserId: 6, VideoId: 2},
	{UserId: 2, VideoId: 3},
	{UserId: 4, VideoId: 3},
	{UserId: 6, VideoId: 3},
	{UserId: 1, VideoId: 4},
	{UserId: 2, VideoId: 4},
	{UserId: 3, VideoId: 4},
	{UserId: 5, VideoId: 4},
	{UserId: 6, VideoId: 4},
	{UserId: 1, VideoId: 5},
	{UserId: 6, VideoId: 5},
	{UserId: 3, VideoId: 5},
	{UserId: 4, VideoId: 5},
	{UserId: 2, VideoId: 6},
	{UserId: 1, VideoId: 6},
	{UserId: 3, VideoId: 6},
	{UserId: 4, VideoId: 6},
	{UserId: 6, VideoId: 7},
	{UserId: 1, VideoId: 7},
	{UserId: 2, VideoId: 7},
	{UserId: 2, VideoId: 8},
	{UserId: 3, VideoId: 8},
	{UserId: 5, VideoId: 8},
	{UserId: 3, VideoId: 9},
	{UserId: 5, VideoId: 9},
	{UserId: 1, VideoId: 10},
	{UserId: 2, VideoId: 10},
	{UserId: 3, VideoId: 10},
	{UserId: 4, VideoId: 10},
	{UserId: 5, VideoId: 10},
	{UserId: 6, VideoId: 10},
}

func initFavorites(db *gorm.DB) {
	var count int64
	db.Model(&models.Favorite{}).Count(&count)
	if count == 0 {
		for _, favorite := range initialFavorites {
			db.Create(&favorite)
		}
	}
}

// ========== 评论表-样例数据 ==========

var initialComments = []models.Comment{
	{UserId: 2, VideoId: 1, CreatedAt: time.Now(), Content: "好啦好啦，都在歌单里啦"},
	{UserId: 4, VideoId: 1, CreatedAt: time.Now(), Content: "爷青回！！！"},
	{UserId: 6, VideoId: 1, CreatedAt: time.Now(), Content: "好活，当赏"},
	{UserId: 5, VideoId: 2, CreatedAt: time.Now(), Content: "这竟然是你唱的"},
	{UserId: 6, VideoId: 2, CreatedAt: time.Now(), Content: "这个戏腔绝了"},
	{UserId: 6, VideoId: 2, CreatedAt: time.Now(), Content: "兰亭序：“而我独缺你一生的了解”\n七里香：“你是我唯一想要的了解”"},
	{UserId: 2, VideoId: 3, CreatedAt: time.Now(), Content: "“不要因为廉价的新鲜感放弃长久的陪伴”"},
	{UserId: 6, VideoId: 3, CreatedAt: time.Now(), Content: "幸好思念无声"},
	{UserId: 2, VideoId: 4, CreatedAt: time.Now(), Content: "说走就走的旅行..."},
	{UserId: 1, VideoId: 4, CreatedAt: time.Now(), Content: "这首歌好像在哪里听过！"},
	{UserId: 5, VideoId: 4, CreatedAt: time.Now(), Content: "好久没出去玩了，走起"},
	{UserId: 3, VideoId: 5, CreatedAt: time.Now(), Content: "变装那一瞬间好高级"},
	{UserId: 1, VideoId: 5, CreatedAt: time.Now(), Content: "这服装好看！"},
	{UserId: 4, VideoId: 5, CreatedAt: time.Now(), Content: "这种变装是怎么做到的！"},
	{UserId: 1, VideoId: 6, CreatedAt: time.Now(), Content: "因世上的至爱是不计较条件..."},
	{UserId: 2, VideoId: 7, CreatedAt: time.Now(), Content: "年少的梦终将绽放于盛夏，解放之鼓必将响彻整个夏天"},
	{UserId: 6, VideoId: 7, CreatedAt: time.Now(), Content: "啧，怎么说呢......"},
	{UserId: 1, VideoId: 7, CreatedAt: time.Now(), Content: "好看"},
	{UserId: 3, VideoId: 8, CreatedAt: time.Now(), Content: "落霞与孤鹜齐飞，秋水共长天一色"},
	{UserId: 5, VideoId: 8, CreatedAt: time.Now(), Content: "为天地立心，为生民立命，为往圣继绝学，为万世开太平"},
	{UserId: 5, VideoId: 9, CreatedAt: time.Now(), Content: "夕阳洒在世界的尽头"},
	{UserId: 1, VideoId: 10, CreatedAt: time.Now(), Content: "运动孩多多少少有点不合群"},
	{UserId: 4, VideoId: 10, CreatedAt: time.Now(), Content: "这圆画的好圆"},
	{UserId: 3, VideoId: 10, CreatedAt: time.Now(), Content: "和我体育老师教的好像一样，但又不怎么一样"},
	{UserId: 2, VideoId: 10, CreatedAt: time.Now(), Content: "行云流水"},
}

func initComments(db *gorm.DB) {
	var count int64
	db.Model(&models.Comment{}).Count(&count)
	if count == 0 {
		for _, comment := range initialComments {
			db.Create(&comment)
		}
	}
}

// ========== 关注关系表-样例数据 ==========

var initialRelations = []models.Relation{
	{FromUserId: 1, ToUserId: 4, Cancel: 0},
	{FromUserId: 1, ToUserId: 5, Cancel: 0},
	{FromUserId: 2, ToUserId: 1, Cancel: 0},
	{FromUserId: 2, ToUserId: 3, Cancel: 0},
	{FromUserId: 2, ToUserId: 6, Cancel: 0},
	{FromUserId: 3, ToUserId: 1, Cancel: 0},
	{FromUserId: 3, ToUserId: 5, Cancel: 0},
	{FromUserId: 3, ToUserId: 6, Cancel: 0},
	{FromUserId: 3, ToUserId: 4, Cancel: 0},
	{FromUserId: 4, ToUserId: 1, Cancel: 0},
	{FromUserId: 4, ToUserId: 3, Cancel: 0},
	{FromUserId: 4, ToUserId: 2, Cancel: 0},
	{FromUserId: 4, ToUserId: 5, Cancel: 0},
	{FromUserId: 5, ToUserId: 1, Cancel: 0},
	{FromUserId: 5, ToUserId: 4, Cancel: 0},
	{FromUserId: 5, ToUserId: 6, Cancel: 0},
	{FromUserId: 6, ToUserId: 1, Cancel: 0},
	{FromUserId: 6, ToUserId: 3, Cancel: 0},
	{FromUserId: 6, ToUserId: 2, Cancel: 0},
	{FromUserId: 6, ToUserId: 4, Cancel: 0},
}

func initRelations(db *gorm.DB) {
	var count int64
	db.Model(&models.Relation{}).Count(&count)
	if count == 0 {
		for _, relation := range initialRelations {
			db.Create(&relation)
		}
	}
}

// ========== 聊天记录表-样例数据 ==========

var initialMessages = []models.Message{
	// 用户一 和 用户二 的初始对话
	{FromUserID: 1, ToUserID: 2, CreateTime: time.Now(), Content: "你会GO吗"},
	{FromUserID: 2, ToUserID: 1, CreateTime: time.Now(), Content: "会一点，怎么了"},
	{FromUserID: 1, ToUserID: 2, CreateTime: time.Now(), Content: "最近做Web应用，有点挑战。"},
	{FromUserID: 2, ToUserID: 1, CreateTime: time.Now(), Content: "嗯，Go在Web开发不错。遇到啥挑战了"},
	{FromUserID: 1, ToUserID: 2, CreateTime: time.Now(), Content: "处理并发和性能，用goroutines处理多任务，但调度和同步有问题"},
	{FromUserID: 2, ToUserID: 1, CreateTime: time.Now(), Content: "那你得注意竞态和内存错误"},
	{FromUserID: 1, ToUserID: 2, CreateTime: time.Now(), Content: "对，我用channels同步数据，但偶尔出bug，很纠结"},
	{FromUserID: 2, ToUserID: 1, CreateTime: time.Now(), Content: "可能是goroutine间通信问题，检查下channel用法，避免死锁"},
	{FromUserID: 1, ToUserID: 2, CreateTime: time.Now(), Content: "我去看一下"},
	{FromUserID: 2, ToUserID: 1, CreateTime: time.Now(), Content: "嗯，有问题再来讨论"},
	{FromUserID: 1, ToUserID: 2, CreateTime: time.Now(), Content: "好！"},
	// 用户二 和 用户三 的初始对话
	{FromUserID: 3, ToUserID: 2, CreateTime: time.Now(), Content: "周末有安排嘛"},
	{FromUserID: 2, ToUserID: 3, CreateTime: time.Now(), Content: "没，怎么说"},
	{FromUserID: 3, ToUserID: 2, CreateTime: time.Now(), Content: "打球哩，走起"},
	{FromUserID: 2, ToUserID: 3, CreateTime: time.Now(), Content: "几点"},
	{FromUserID: 3, ToUserID: 2, CreateTime: time.Now(), Content: "六点吧"},
	{FromUserID: 2, ToUserID: 3, CreateTime: time.Now(), Content: "行，你叫上他们，人多点"},
	{FromUserID: 3, ToUserID: 2, CreateTime: time.Now(), Content: "已经在约了"},
	{FromUserID: 2, ToUserID: 3, CreateTime: time.Now(), Content: "OK"},
}

func initMessages(db *gorm.DB) {
	var count int64
	db.Model(&models.Message{}).Count(&count)
	if count == 0 {
		for _, message := range initialMessages {
			db.Create(&message)
		}
	}
}

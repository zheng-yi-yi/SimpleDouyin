package config

import (
	"time"

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
	// ========== 初始化用户表 ==========
	var count_User int64
	db.Model(&models.User{}).Count(&count_User)
	if count_User == 0 {
		user1 := models.User{
			UserName:        "User1",          // 用户名
			PassWord:        "123456",         // 密码
			Status:          1,                // 用户状态
			FollowCount:     2,                // 关注总数（关注User2、User3）
			FollowerCount:   2,                // 粉丝总数（User2、User3）
			FavoriteCount:   3,                // 喜欢数（video: 2、4、5）
			Avatar:          User1_avater,     // 用户头像
			BackgroundImage: User1_background, // 用户个人页顶部大图
			Signature:       User1_signature,  // 个人简介
			TotalFavorited:  "6",              // 获赞数量（共获6个赞）
			WorkCount:       3,                // 作品数（发布了三个作品：1、3、6）
		}
		db.Create(&user1)
		user2 := models.User{
			UserName:        "User2",          // 用户名
			PassWord:        "123456",         // 密码
			Status:          1,                // 用户状态
			FollowCount:     2,                // 关注总数（关注User1、User3）
			FollowerCount:   1,                // 粉丝总数（user3）
			FavoriteCount:   4,                // 喜欢数（video: 1、3、5、6）
			Avatar:          User2_avater,     // 用户头像
			BackgroundImage: User2_background, // 用户个人页顶部大图
			Signature:       User2_signature,  // 个人简介
			TotalFavorited:  "4",              // 获赞数量（共获4个赞）
			WorkCount:       2,                // 作品数（发布了两个作品:2、4）
		}
		db.Create(&user2)
		user3 := models.User{
			UserName:        "User3",          // 用户名
			PassWord:        "123456",         // 密码
			Status:          1,                // 用户状态
			FollowCount:     2,                // 关注总数（关注User1、User2）
			FollowerCount:   0,                // 粉丝总数（没有人关注）
			FavoriteCount:   5,                // 喜欢数（video: 1、2、3、4、6）
			Avatar:          User3_avater,     // 用户头像
			BackgroundImage: User3_background, // 用户个人页顶部大图
			Signature:       User3_signature,  // 个人简介
			TotalFavorited:  "2",              // 获赞数量（共获2个赞）
			WorkCount:       3,                // 作品数（发布了一个作品：5）
		}
		db.Create(&user3)
	}
	// ========== 初始化视频表 ==========
	var count_Video int64
	db.Model(&models.Video{}).Count(&count_Video)
	if count_Video == 0 {
		videos := []*models.Video{
			{UserId: 1, PlayUrl: "videos/1.mp4", CoverUrl: "images/1.jpg", FavoriteCount: 2, CommentCount: 2, Description: "邀你观看2023稀土开发者大会!11场论坛，60位技术大咖，涵盖AIGC与大模型、大前端、音视频等前沿技术资讯"},
			{UserId: 2, PlayUrl: "videos/2.mp4", CoverUrl: "images/2.jpg", FavoriteCount: 2, CommentCount: 2, Description: "那些仅凭半句就封神的诗句"},
			{UserId: 1, PlayUrl: "videos/3.mp4", CoverUrl: "images/3.jpg", FavoriteCount: 2, CommentCount: 2, Description: "登昆仑兮食玉英，与天地兮同寿与日月兮齐光。#汉服之美在华夏"},
			{UserId: 2, PlayUrl: "videos/4.mp4", CoverUrl: "images/4.jpg", FavoriteCount: 2, CommentCount: 2, Description: "落日沉溺于橘色的海，晚风沦陷于赤城的爱"},
			{UserId: 3, PlayUrl: "videos/5.mp4", CoverUrl: "images/5.jpg", FavoriteCount: 2, CommentCount: 2, Description: "五档 启动！ #太阳神尼卡登场"},
			{UserId: 1, PlayUrl: "videos/6.mp4", CoverUrl: "images/6.jpg", FavoriteCount: 2, CommentCount: 2, Description: "好像做了一场短暂的山水梦"},
		}
		db.Create(&videos)
	}
	// ========== 初始化点赞表 ==========
	var count_Favorite int64
	db.Model(&models.Favorite{}).Count(&count_Favorite)
	if count_Favorite == 0 {
		favorite := []*models.Favorite{
			{UserId: 2, VideoId: 1, Status: 1},
			{UserId: 3, VideoId: 1, Status: 1},
			{UserId: 1, VideoId: 2, Status: 1},
			{UserId: 3, VideoId: 2, Status: 1},
			{UserId: 2, VideoId: 3, Status: 1},
			{UserId: 3, VideoId: 3, Status: 1},
			{UserId: 1, VideoId: 4, Status: 1},
			{UserId: 3, VideoId: 4, Status: 1},
			{UserId: 1, VideoId: 5, Status: 1},
			{UserId: 2, VideoId: 5, Status: 1},
			{UserId: 2, VideoId: 6, Status: 1},
			{UserId: 3, VideoId: 6, Status: 1},
		}
		db.Create(&favorite)
	}
	// ========== 初始化评论表 ==========
	var count_Comment int64
	db.Model(&models.Comment{}).Count(&count_Comment)
	if count_Comment == 0 {
		comment := []*models.Comment{
			{UserId: 2, VideoId: 1, CreatedAt: time.Now(), Cancel: 0, Content: "来学习了"},
			{UserId: 3, VideoId: 1, CreatedAt: time.Now(), Cancel: 0, Content: "真不错！"},
			{UserId: 1, VideoId: 2, CreatedAt: time.Now(), Cancel: 0, Content: "为天地立心，为生民立命，为往圣继绝学，为万世开太平"},
			{UserId: 3, VideoId: 2, CreatedAt: time.Now(), Cancel: 0, Content: "落霞与孤鹜齐飞，秋水共长天一色。"},
			{UserId: 2, VideoId: 3, CreatedAt: time.Now(), Cancel: 0, Content: "变装那一瞬间好高级"},
			{UserId: 3, VideoId: 3, CreatedAt: time.Now(), Cancel: 0, Content: "汉服美"},
			{UserId: 1, VideoId: 4, CreatedAt: time.Now(), Cancel: 0, Content: "我对着日落许愿，希望你永远快乐"},
			{UserId: 3, VideoId: 4, CreatedAt: time.Now(), Cancel: 0, Content: "夕阳洒在世界的尽头"},
			{UserId: 1, VideoId: 5, CreatedAt: time.Now(), Cancel: 0, Content: "年少的梦终将绽放于盛夏，解放之鼓必将响彻整个夏天"},
			{UserId: 2, VideoId: 5, CreatedAt: time.Now(), Cancel: 0, Content: "啧，怎么说呢......"},
			{UserId: 2, VideoId: 6, CreatedAt: time.Now(), Cancel: 0, Content: "说走就走的旅行..."},
			{UserId: 3, VideoId: 6, CreatedAt: time.Now(), Cancel: 0, Content: "这首歌好像在哪里听过！"},
		}
		db.Create(&comment)
	}
	// ========== 初始化关注关系表 ==========
	var count_Relation int64
	db.Model(&models.Relation{}).Count(&count_Relation)
	if count_Relation == 0 {
		relation := []*models.Relation{
			{FromUserId: 1, ToUserId: 2, Cancel: 0},
			{FromUserId: 1, ToUserId: 3, Cancel: 0},
			{FromUserId: 2, ToUserId: 1, Cancel: 0},
			{FromUserId: 2, ToUserId: 3, Cancel: 0},
			{FromUserId: 3, ToUserId: 1, Cancel: 0},
			{FromUserId: 3, ToUserId: 2, Cancel: 0},
		}
		db.Create(&relation)
	}
	// ========== 初始化聊天信息表 ==========
	var count_message int64
	db.Model(&models.Message{}).Count(&count_message)
	if count_message == 0 {
		message := []*models.Message{
			{FromUserID: 1, ToUserID: 2, CreateTime: time.Now(), Content: "Hey! Just got back from the gym. Crushed my leg day workout."},
			{FromUserID: 2, ToUserID: 1, CreateTime: time.Now(), Content: "Nice one! I wish I had that kind of dedication. How was your workout?"},
			{FromUserID: 1, ToUserID: 2, CreateTime: time.Now(), Content: "It was intense! Squats, lunges, and some deadlifts. My legs are already feeling like jelly. How's your day going?"},
			{FromUserID: 2, ToUserID: 1, CreateTime: time.Now(), Content: "Haha, I can only imagine. My day's been pretty good, actually. I managed to hit my word count goal for the day. Writing flow was on point."},
			{FromUserID: 1, ToUserID: 2, CreateTime: time.Now(), Content: "That's awesome! Your writing discipline always impresses me. What are you working on lately?"},
			{FromUserID: 2, ToUserID: 1, CreateTime: time.Now(), Content: "Thanks! Right now, I'm deep into a mystery novel. Trying to weave in some unexpected twists to keep the readers hooked. By the way, got any post-workout meal suggestions? I need some fuel."},
			{FromUserID: 1, ToUserID: 2, CreateTime: time.Now(), Content: "Mystery novel sounds intriguing! As for post-workout meals, how about a grilled chicken salad with lots of veggies and a light vinaigrette? Balanced and delicious."},
			{FromUserID: 2, ToUserID: 1, CreateTime: time.Now(), Content: "Sounds delicious! I'll give that a shot. By the way, I'm attending a literary event this weekend. Care to join and share some of your fitness wisdom during the open mic session?"},
			{FromUserID: 1, ToUserID: 2, CreateTime: time.Now(), Content: "I appreciate the invite! It sounds fun, but I'll actually be out of town for a fitness workshop. Maybe next time though. Break a leg at the event!"},
			{FromUserID: 2, ToUserID: 1, CreateTime: time.Now(), Content: "No problem at all, have a great time at the workshop! We'll catch up when you're back. Enjoy those gains and take lots of notes."},
			{FromUserID: 1, ToUserID: 2, CreateTime: time.Now(), Content: "Will do! And you, keep those creative juices flowing. Looking forward to hearing all about the event."},
			{FromUserID: 2, ToUserID: 1, CreateTime: time.Now(), Content: "Absolutely, I'll make sure to share the highlights. Have a fantastic workout and a productive workshop! Talk to you soon."},
			{FromUserID: 2, ToUserID: 3, CreateTime: time.Now(), Content: "Hey! Just hit a major milestone in my novel today - finished the first draft!"},
			{FromUserID: 3, ToUserID: 2, CreateTime: time.Now(), Content: "That's incredible news! Congratulations! I can only imagine the feeling of completing a first draft. How do you celebrate such writing victories?"},
			{FromUserID: 2, ToUserID: 3, CreateTime: time.Now(), Content: "Thanks! It's definitely a surreal feeling. To celebrate, I usually take a long walk in the park and let my thoughts wander. Speaking of nature, have you been on any exciting outdoor adventures lately?"},
			{FromUserID: 3, ToUserID: 2, CreateTime: time.Now(), Content: "Walking in the park sounds like a perfect celebration. And yes, I actually went camping by the lakeside last weekend. Woke up to the most stunning sunrise over the water. By the way, any tips for overcoming writer's block? I could use some advice."},
			{FromUserID: 2, ToUserID: 3, CreateTime: time.Now(), Content: "Camping by the lakeside sounds like a dream! For writer's block, I usually switch up my writing environment or do a bit of free writing to get the ideas flowing. Nature walks help too. How about planning a nature retreat to recharge your creativity?"},
			{FromUserID: 3, ToUserID: 2, CreateTime: time.Now(), Content: "Those are great suggestions, thank you! A nature retreat sounds amazing right now. Maybe I'll plan a weekend getaway soon. Also, there's a nature photography exhibit coming up. Would you be interested in attending and sharing your writing journey during the panel discussion?"},
			{FromUserID: 2, ToUserID: 3, CreateTime: time.Now(), Content: "A nature photography exhibit sounds intriguing! Unfortunately, I'm swamped with revisions at the moment. But I'd love to catch up afterward and hear all about the event. Best of luck with the panel discussion!"},
			{FromUserID: 3, ToUserID: 2, CreateTime: time.Now(), Content: "Totally understand. I'll make sure to fill you in on all the details. Keep rocking those revisions and let's definitely plan that catch-up. Oh, and any book recommendations for my upcoming solo camping trip?"},
			{FromUserID: 2, ToUserID: 3, CreateTime: time.Now(), Content: "Thanks! For your camping trip, I'd recommend a classic like `Walden` by Henry David Thoreau. It's perfect for immersing yourself in nature's beauty. Have an incredible trip and get inspired!"},
			{FromUserID: 3, ToUserID: 2, CreateTime: time.Now(), Content: "`Walden` it is! Thanks for the recommendation. I'm looking forward to some quality nature reading. Enjoy your revisions and we'll chat soon. Happy writing!"},
			{FromUserID: 2, ToUserID: 3, CreateTime: time.Now(), Content: "Thank you! Enjoy your camping trip and take in all the natural wonders. Looking forward to our next chat. Take care!"},
		}
		db.Create(&message)
	}
	return nil
}

package config

const (
	// 资源路径
	SERVER_RESOURCES     = "http://" + LOCAL_IP_ADDRESS + ":8080/public/"
	DEFAULT_AVATAR_URL   = SERVER_RESOURCES + "Initdata/avatar/"
	DEFAULT_BG_IMAGE_URL = SERVER_RESOURCES + "Initdata/background/"

	// 默认用户信息
	DEFAULT_USER_AVATAR_URL   = DEFAULT_AVATAR_URL + "default.png"   // 默认头像地址
	DEFAULT_USER_BG_IMAGE_URL = DEFAULT_BG_IMAGE_URL + "default.png" // 默认背景图地址
	DEFAULT_USER_BIO          = "这个人很懒，什么也没有留下......"                // 默认简介内容

	// 用户 1
	USER1_AVATAR_URL           = DEFAULT_AVATAR_URL + "1.png"
	USER1_BACKGROUND_IMAGE_URL = DEFAULT_BG_IMAGE_URL + "1.png"
	USER1_PROFILE_DESCRIPTION  = "一个积极追求健康生活方式的健身爱好者"

	// 用户 2
	USER2_AVATAR_URL           = DEFAULT_AVATAR_URL + "2.png"
	USER2_BACKGROUND_IMAGE_URL = DEFAULT_BG_IMAGE_URL + "2.png"
	USER2_PROFILE_DESCRIPTION  = "编织奇幻世界，让每个故事都成为一场无限想象的冒险。"

	// 用户 3
	USER3_AVATAR_URL           = DEFAULT_AVATAR_URL + "3.png"
	USER3_BACKGROUND_IMAGE_URL = DEFAULT_BG_IMAGE_URL + "3.png"
	USER3_PROFILE_DESCRIPTION  = "摄影的魅力在于记录历史，让历史定格在一瞬间，让瞬间成为永恒。"

	// 用户 4
	USER4_AVATAR_URL           = DEFAULT_AVATAR_URL + "4.png"
	USER4_BACKGROUND_IMAGE_URL = DEFAULT_BG_IMAGE_URL + "4.png"
	USER4_PROFILE_DESCRIPTION  = "探索自然奥秘、追求精神自由，驰骋大海，是冲浪爱好者的“浪尖之舞”。"

	// 用户 5
	USER5_AVATAR_URL           = DEFAULT_AVATAR_URL + "5.png"
	USER5_BACKGROUND_IMAGE_URL = DEFAULT_BG_IMAGE_URL + "5.png"
	USER5_PROFILE_DESCRIPTION  = "艺术是一种审美的意识形态，其反映本质是审美的而不是科学的。"

	// 用户 6
	USER6_AVATAR_URL           = DEFAULT_AVATAR_URL + "6.png"
	USER6_BACKGROUND_IMAGE_URL = DEFAULT_BG_IMAGE_URL + "6.png"
	USER6_PROFILE_DESCRIPTION  = "自然爱好者，通过镜头捕捉大自然的美丽，分享独特的视角和故事。"
)

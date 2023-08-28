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
	USER1_PROFILE_DESCRIPTION  = "明确地爱，直接地厌恶，真诚地喜欢，站在太阳下的坦荡，大声无愧地称赞自己"

	// 用户 2
	USER2_AVATAR_URL           = DEFAULT_AVATAR_URL + "2.png"
	USER2_BACKGROUND_IMAGE_URL = DEFAULT_BG_IMAGE_URL + "2.png"
	USER2_PROFILE_DESCRIPTION  = "发呆业务爱好者"

	// 用户 3
	USER3_AVATAR_URL           = DEFAULT_AVATAR_URL + "3.png"
	USER3_BACKGROUND_IMAGE_URL = DEFAULT_BG_IMAGE_URL + "3.png"
	USER3_PROFILE_DESCRIPTION  = "这里介绍不了我"

	// 用户 4
	USER4_AVATAR_URL           = DEFAULT_AVATAR_URL + "4.png"
	USER4_BACKGROUND_IMAGE_URL = DEFAULT_BG_IMAGE_URL + "4.png"
	USER4_PROFILE_DESCRIPTION  = "有时候词不达意 但我很高兴遇见你"

	// 用户 5
	USER5_AVATAR_URL           = DEFAULT_AVATAR_URL + "5.png"
	USER5_BACKGROUND_IMAGE_URL = DEFAULT_BG_IMAGE_URL + "5.png"
	USER5_PROFILE_DESCRIPTION  = "周末就是将生活调回自己喜欢的频道"

	// 用户 6
	USER6_AVATAR_URL           = DEFAULT_AVATAR_URL + "6.png"
	USER6_BACKGROUND_IMAGE_URL = DEFAULT_BG_IMAGE_URL + "6.png"
	USER6_PROFILE_DESCRIPTION  = "麻麻说简介太长就会有笨蛋跟着念"
)

package test

import (
	"net/http"
	"testing"
)

// 测试 "点赞" 模块
func TestFavorite(t *testing.T) {
	e := newExpect(t)

	// 发送 GET 请求来获取抖音动态信息
	feedResp := e.GET("/douyin/feed/").Expect().Status(http.StatusOK).JSON().Object()
	feedResp.Value("status_code").Number().IsEqual(0)   // 确保响应状态码为 0，表示成功
	feedResp.Value("video_list").Array().Length().Gt(0) // 确保视频列表不为空

	// 获取第一个视频的 ID
	firstVideo := feedResp.Value("video_list").Array().Value(0).Object()
	videoId := firstVideo.Value("id").Number().Raw()

	// 随机得到一个用户名以及密码
	var favorite_username = generateRandomUsername()
	var favorite_password = generateRandomPassword()

	// 获取测试用户的用户 ID 和 token
	userId, token := getTestUserIdAndToken(favorite_username, favorite_password, e)

	// 发送 POST 请求来执行点赞操作
	favoriteResp := e.POST("/douyin/favorite/action/").
		WithQuery("token", token).WithQuery("video_id", videoId).WithQuery("action_type", 1).
		WithFormField("token", token).WithFormField("video_id", videoId).WithFormField("action_type", 1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	favoriteResp.Value("status_code").Number().IsEqual(0) // 确保收藏操作成功

	// 发送 GET 请求来获取用户的点赞列表
	favoriteListResp := e.GET("/douyin/favorite/list/").
		WithQuery("token", token).WithQuery("user_id", userId).
		WithFormField("token", token).WithFormField("user_id", userId).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	favoriteListResp.Value("status_code").String().IsEqual("0") // 确保获取收藏列表成功

	// 遍历收藏列表中的每个视频并进行验证
	for _, element := range favoriteListResp.Value("video_list").Array().Iter() {
		video := element.Object()
		video.ContainsKey("id")                      // 确保视频对象包含 ID
		video.ContainsKey("author")                  // 确保视频对象包含作者信息
		video.Value("play_url").String().NotEmpty()  // 确保播放链接不为空
		video.Value("cover_url").String().NotEmpty() // 确保封面链接不为空
	}
}

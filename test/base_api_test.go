package test

import (
	"net/http"
	"testing"
)

// 测试 “视频流” 功能
func TestFeed(t *testing.T) {
	e := newExpect(t)
	feedResp := e.GET("/douyin/feed/").Expect().Status(http.StatusOK).JSON().Object()
	feedResp.Value("status_code").Number().IsEqual(0)
	feedResp.Value("video_list").Array().Length().Gt(0)
	for _, element := range feedResp.Value("video_list").Array().Iter() {
		video := element.Object()
		video.ContainsKey("id")
		video.ContainsKey("author")
		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()
	}
}

// 测试 "用户" 模块
func TestUserAction(t *testing.T) {

	e := newExpect(t)

	// 随机得到一个用户名以及密码
	var registerName = generateRandomUsername()
	var registerPwd = generateRandomPassword()

	// 注册用户并验证响应
	registerResp := e.POST("/douyin/user/register/").
		WithQuery("username", registerName).
		WithQuery("password", registerPwd).
		WithFormField("username", registerName).
		WithFormField("password", registerPwd).
		Expect().
		// 确保HTTP状态码为200
		Status(http.StatusOK).
		// 解析响应为JSON对象
		JSON().Object()
	// 验证注册响应中的字段是否符合预期
	registerResp.Value("status_code").Number().IsEqual(0)
	registerResp.Value("user_id").Number().Gt(0)
	registerResp.Value("token").String().Length().Gt(0)

	// 登录用户并验证响应
	loginResp := e.POST("/douyin/user/login/").
		WithQuery("username", registerName).
		WithQuery("password", registerPwd).
		WithFormField("username", registerName).
		WithFormField("password", registerPwd).
		Expect().
		// 确保HTTP状态码为200
		Status(http.StatusOK).
		// 解析响应为JSON对象
		JSON().Object()
	// 验证登录响应中的字段
	loginResp.Value("status_code").Number().IsEqual(0)
	loginResp.Value("user_id").Number().Gt(0)
	loginResp.Value("token").String().Length().Gt(0)

	// 获取登录后的 token 和 用户 id。
	token := loginResp.Value("token").String().Raw()
	UserId := loginResp.Value("user_id").Number().Raw()

	// 获取用户信息并验证响应
	userResp := e.GET("/douyin/user/").
		WithQuery("token", token).
		WithQuery("user_id", UserId).
		Expect().
		// 确保HTTP状态码为200
		Status(http.StatusOK).
		// 解析响应为JSON对象
		JSON().Object()
	// 验证用户信息响应中的字段
	userResp.Value("status_code").Number().IsEqual(0)
	userInfo := userResp.Value("user").Object()
	userInfo.NotEmpty()
	userInfo.Value("id").Number().Gt(0)
	userInfo.Value("name").String().Length().Gt(0)
}

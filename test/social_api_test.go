package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试 “关系” 模块
func TestRelation(t *testing.T) {

	e := newExpect(t)

	// 获取测试用户A的用户 ID 和 token
	userIdA, tokenA := getTestUserIdAndToken("RelationTestA", "114589", e)
	// 获取测试用户B的用户 ID 和令牌
	userIdB, tokenB := getTestUserIdAndToken("RelationTestB", "223657", e)

	// 发送 POST 请求来执行关注操作，将用户 A 关注用户 B
	relationResp := e.POST("/douyin/relation/action/").
		WithQuery("token", tokenA).WithQuery("to_user_id", userIdB).WithQuery("action_type", 1).
		WithFormField("token", tokenA).WithFormField("to_user_id", userIdB).WithFormField("action_type", 1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	relationResp.Value("status_code").Number().IsEqual(0) // 确保关注操作成功

	// 发送 GET 请求来获取用户 A 的关注列表
	followListResp := e.GET("/douyin/relation/follow/list/").
		WithQuery("token", tokenA).WithQuery("user_id", userIdA).
		WithFormField("token", tokenA).WithFormField("user_id", userIdA).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	followListResp.Value("status_code").String().IsEqual("0") // 确保获取关注列表成功

	// 检查关注列表中是否包含测试用户 B
	containTestUserB := false
	for _, element := range followListResp.Value("user_list").Array().Iter() {
		user := element.Object()
		user.ContainsKey("id") // 确保用户对象包含 ID
		if int(user.Value("id").Number().Raw()) == userIdB {
			containTestUserB = true
		}
	}
	// 使用断言确保测试用户 B 在关注列表中
	assert.True(t, containTestUserB, "关注测试用户失败")

	// 发送 GET 请求来获取用户 B 的粉丝列表
	followerListResp := e.GET("/douyin/relation/follower/list/").
		WithQuery("token", tokenB).WithQuery("user_id", userIdB).
		WithFormField("token", tokenB).WithFormField("user_id", userIdB).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	followerListResp.Value("status_code").String().IsEqual("0") // 确保获取粉丝列表成功

	// 检查粉丝列表中是否包含测试用户 A
	containTestUserA := false
	for _, element := range followerListResp.Value("user_list").Array().Iter() {
		user := element.Object()
		user.ContainsKey("id") // 确保用户对象包含 ID
		if int(user.Value("id").Number().Raw()) == userIdA {
			containTestUserA = true
		}
	}
	// 使用断言确保测试用户 A 在粉丝列表中
	assert.True(t, containTestUserA, "关注者测试用户失败")
}

// 测试 “消息” 模块
func TestChat(t *testing.T) {

	e := newExpect(t)

	// 获取测试用户A的用户 ID 和 token
	userIdA, tokenA := getTestUserIdAndToken("ChatTestA", "881246", e)
	// 获取测试用户B的用户 ID 和令牌
	userIdB, tokenB := getTestUserIdAndToken("ChatTestB", "958831", e)

	// 发送 POST 请求来发送消息给用户 B
	messageResp := e.POST("/douyin/message/action/").
		WithQuery("token", tokenA).WithQuery("to_user_id", userIdB).WithQuery("action_type", 1).WithQuery("content", "Send to UserB").
		WithFormField("token", tokenA).WithFormField("to_user_id", userIdB).WithFormField("action_type", 1).WithQuery("content", "Send to UserB").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	messageResp.Value("status_code").Number().IsEqual(0) // 确保发送消息操作成功

	// 发送 GET 请求来获取用户 A 和用户 B 之间的聊天消息列表
	chatResp := e.GET("/douyin/message/chat/").
		WithQuery("token", tokenA).WithQuery("to_user_id", userIdB).
		WithFormField("token", tokenA).WithFormField("to_user_id", userIdB).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	chatResp.Value("status_code").String().IsEqual("0") // 确保获取聊天消息列表成功
	chatResp.Value("message_list").Array().NotEmpty()   // 确保聊天消息列表不为空

	// 发送 GET 请求来获取用户 B 和用户 A 之间的聊天消息列表
	chatResp = e.GET("/douyin/message/chat/").
		WithQuery("token", tokenB).WithQuery("to_user_id", userIdA).
		WithFormField("token", tokenB).WithFormField("to_user_id", userIdA).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	chatResp.Value("status_code").String().IsEqual("0") // 确保获取聊天消息列表成功
	chatResp.Value("message_list").Array().NotEmpty()   // 确保聊天消息列表不为空

}

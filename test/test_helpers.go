package test

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
	"github.com/zheng-yi-yi/simpledouyin/config"
)

var serverAddr = "http://" + config.LOCAL_IP_ADDRESS + ":8080" // 服务器地址

// newExpect , 创建并返回一个新的 httpexpect.Expect 实例，用于进行 HTTP 请求和断言
func newExpect(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		Client:   http.DefaultClient,
		BaseURL:  serverAddr,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewCompactPrinter(t),
		},
	})
}

// getTestUserIdAndToken , 获取测试用户的并获取其ID以及token
func getTestUserIdAndToken(username string, password string, e *httpexpect.Expect) (int, string) {
	registerResp := e.POST("/douyin/user/register/").
		WithQuery("username", username).WithQuery("password", password).
		WithFormField("username", username).WithFormField("password", password).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	userId := 0
	token := registerResp.Value("token").String().Raw()

	if len(token) == 0 {
		loginResp := e.POST("/douyin/user/login/").
			WithQuery("username", username).WithQuery("password", password).
			WithFormField("username", username).WithFormField("password", password).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		loginToken := loginResp.Value("token").String()
		loginToken.Length().Gt(0)
		token = loginToken.Raw()

		userId = int(loginResp.Value("user_id").Number().Raw())
	} else {
		userId = int(registerResp.Value("user_id").Number().Raw())
	}
	return userId, token
}

// generateRandomUsername ， 生成随机用户名
func generateRandomUsername() string {
	// 定义可用字符集合
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)
	// 生成随机8个字符的用户名
	username := make([]byte, 8)
	for i := 0; i < 8; i++ {
		username[i] = characters[randGen.Intn(len(characters))]
	}
	return string(username)
}

// generateRandomPassword ， 生成随机六位数密码
func generateRandomPassword() string {
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)
	password := fmt.Sprintf("%06d", randGen.Intn(1000000))
	return password
}

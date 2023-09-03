package test

import (
	"net/http"
	"testing"

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

package rest_test

import (
	"context"
	"testing"

	"github.com/yoas0bi/micro-toolkit/client/rest"
	"github.com/yoas0bi/micro-toolkit/http/response"
	"github.com/yoas0bi/micro-toolkit/logger/zap"
)

func TestClient(t *testing.T) {
	c := rest.NewRESTClient()
	c.SetBaseURL("https://www.baidu.com")
	c.SetBearerTokenAuth("UAoVkI07gDGlfARUTToCA8JW")

	var h string
	resp := make(map[string]any)
	err := c.Group("group1").Group("group2").Get("/getpath").Prefix("pre").Suffix("sub").Param("test", "test01").
		Do(context.Background()).
		Header("Server", &h).
		Into(response.NewData(&resp))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(h)
}

func init() {
	// 设置日志模式
	zap.DevelopmentSetup()
}

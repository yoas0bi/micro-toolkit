package request_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yoas0bi/micro-toolkit/http/request"
)

func TestGetRemoteIpFromHeader(t *testing.T) {
	shoud := assert.New(t)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Forwarded-For", "10.10.10.10")
	ip := request.GetRemoteIP(req)
	shoud.Equal("10.10.10.10", ip)
}

func TestGetRemoteIpFromConn(t *testing.T) {
	shoud := assert.New(t)

	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "10.10.10.10"
	ip := request.GetRemoteIP(req)
	shoud.Equal("10.10.10.10", ip)
}

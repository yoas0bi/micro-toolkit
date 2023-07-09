package router_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yoas0bi/micro-toolkit/http/router"
)

func simpleAuth(h http.Header) (authInfo interface{}, err error) {
	authHeader := h.Get("Authorization")
	return authHeader, nil
}

func TestAutherFunc(t *testing.T) {
	should := require.New(t)

	auther := router.AutherFunc(simpleAuth)
	header := make(http.Header)
	header.Add("Authorization", "ok")
	authInfo, err := auther.Auth(header)
	should.NoError(err)
	should.Equal("ok", authInfo)
}

package router_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yoas0bi/micro-toolkit/http/router"
)

type handler struct{}

func (h *handler) FuncWithStruct(w http.ResponseWriter, r *http.Request) {}
func FuncNoStruct(w http.ResponseWriter, r *http.Request)                {}

func TestGetHandlerFuncNameWithStruct(t *testing.T) {
	should := require.New(t)

	h := new(handler)
	fn := router.GetHandlerFuncName(h.FuncWithStruct)
	should.Equal("FuncWithStruct", fn)
}

func TestGetHandlerFuncNameWithNoStruct(t *testing.T) {
	should := require.New(t)

	fn := router.GetHandlerFuncName(FuncNoStruct)
	should.Equal("FuncNoStruct", fn)
}

package recovery_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yoas0bi/micro-toolkit/http/middleware/recovery"
	"github.com/yoas0bi/micro-toolkit/http/router/httprouter"
	"github.com/yoas0bi/micro-toolkit/logger/zap"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	panic("recovery test")
}

func Test_Recovery(t *testing.T) {
	should := require.New(t)

	router := httprouter.New()

	rm := recovery.NewWithLogger(zap.L())
	router.Use(rm)
	router.Handle("GET", "/", indexHandler)

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://localhost:3000/", nil)
	should.NoError(err)

	router.ServeHTTP(recorder, req)
	should.Equal(recorder.Code, http.StatusInternalServerError)
}

func init() {
	zap.DevelopmentSetup()
	zap.L()
}

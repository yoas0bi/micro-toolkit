package sysutil_test

import (
	"os"
	"runtime"
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/sysutil"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestBasic_usage(t *testing.T) {
	assert.NotEmpty(t, sysutil.BinDir())
	assert.NotEmpty(t, sysutil.BinDir())
	assert.NotEmpty(t, sysutil.BinFile())
}

func TestProcessExists(t *testing.T) {
	if runtime.GOOS != "windows" {
		pid := os.Getpid()
		assert.True(t, sysutil.ProcessExists(pid))
	} else {
		t.Skip("on Windows")
	}
}

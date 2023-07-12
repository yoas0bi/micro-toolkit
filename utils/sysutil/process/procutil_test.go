package process_test

import (
	"os"
	"syscall"
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/sysutil/process"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestProcessExists(t *testing.T) {
	pid := os.Getpid()

	assert.True(t, process.Exists(pid))
}

func TestPID(t *testing.T) {
	assert.True(t, process.PID() > 0)
}

func TestKillByName(t *testing.T) {
	assert.Err(t, process.KillByName("never-never-exist-process", syscall.SIGKILL))
}

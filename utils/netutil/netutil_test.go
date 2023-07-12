package netutil_test

import (
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/netutil"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestInternalIP(t *testing.T) {
	assert.NotEmpty(t, netutil.InternalIPv1())
	assert.NotEmpty(t, netutil.InternalIP())
	assert.NotEmpty(t, netutil.InternalIPv4())
	assert.NotEmpty(t, netutil.InternalIPv6())
	assert.NotEmpty(t, netutil.GetLocalIPs())
}

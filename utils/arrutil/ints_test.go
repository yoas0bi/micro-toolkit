package arrutil_test

import (
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/arrutil"
	"github.com/yoas0bi/micro-toolkit/utils-toolkit/utils/testutil/assert"
)

func TestIntsToString(t *testing.T) {
	assert.Eq(t, "[]", arrutil.IntsToString([]int{}))
	assert.Eq(t, "[1,2,3]", arrutil.IntsToString([]int{1, 2, 3}))
}

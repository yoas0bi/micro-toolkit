package panics_test

import (
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/errorx/panics"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestIsTrue(t *testing.T) {
	assert.Panics(t, func() {
		panics.IsTrue(false)
	})
}

package stdutil_test

import (
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/stdutil"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestGo(t *testing.T) {
	err := stdutil.Go(func() error {
		return nil
	})
	assert.NoErr(t, err)
}

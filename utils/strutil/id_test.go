package strutil_test

import (
	"fmt"
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/strutil"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestMicroTimeID(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := strutil.MicroTimeID()
		fmt.Println(id, "len:", len(id))
		assert.NotEmpty(t, id)
	}
}

func TestMicroTimeHexID(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := strutil.MicroTimeHexID()
		fmt.Println(id, "len:", len(id))
		assert.NotEmpty(t, id)
	}
}

func TestDatetimeNo(t *testing.T) {
	for i := 0; i < 10; i++ {
		no := strutil.DatetimeNo("test")
		fmt.Println(no, "len:", len(no))
		assert.NotEmpty(t, no)
	}
}

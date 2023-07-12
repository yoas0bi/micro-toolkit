package stdio_test

import (
	"bytes"
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/stdio"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestNewWriteWrapper(t *testing.T) {
	buf := new(bytes.Buffer)

	w := stdio.NewWriteWrapper(buf)
	_, err := w.WriteString("inhere")
	assert.NoErr(t, err)
	assert.Eq(t, "inhere", w.String())

	err = w.WriteByte(',')
	assert.NoErr(t, err)

	_, err = w.Write([]byte("hi."))
	assert.NoErr(t, err)
	assert.Eq(t, "inhere,hi.", w.String())

	_, err = w.Writef(" ok, %s.", "tom")
	assert.NoErr(t, err)
	assert.Eq(t, "inhere,hi. ok, tom.", w.String())
}

package fakeobj_test

import (
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/fakeobj"
)

func TestNewWriter(t *testing.T) {
	tw := fakeobj.NewBuffer()
	_, err := tw.Write([]byte("hello"))
	assert.NoErr(t, err)
	assert.Eq(t, "hello", tw.String())
	assert.NoErr(t, tw.Flush())
	assert.Eq(t, "", tw.String())
	assert.NoErr(t, tw.Close())

	// write string
	_, err = tw.WriteString("hello")
	assert.NoErr(t, err)
	assert.Eq(t, "hello", tw.ResetGet())

	tw.SetErrOnWrite()
	_, err = tw.Write([]byte("hello"))
	assert.Err(t, err)
	assert.Eq(t, "", tw.String())

	tw.SetErrOnFlush()
	assert.Err(t, tw.Flush())

	tw.SetErrOnClose()
	assert.Err(t, tw.Close())
}

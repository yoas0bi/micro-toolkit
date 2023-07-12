package clipboard_test

import (
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/sysutil/clipboard"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestGetReaderBin(t *testing.T) {
	assert.NotEmpty(t, clipboard.GetReaderBin())
	assert.NotEmpty(t, clipboard.GetWriterBin())
	assert.NotEmpty(t, clipboard.Std())
}

func TestClipboard_read_write(t *testing.T) {
	err := clipboard.WriteString("")
	assert.ErrMsg(t, err, "clipboard: empty contents for write")

	if !clipboard.Available() {
		assert.False(t, clipboard.Available())
		t.Skipf("skip test on program '%s' not found", clipboard.GetReaderBin())
		return
	}

	err = clipboard.Reset()
	assert.NoErr(t, err)

	str, err := clipboard.ReadString()
	assert.NoErr(t, err)
	assert.Empty(t, str)

	src := "hello, this is clipboard"
	err = clipboard.WriteString(src)
	assert.NoErr(t, err)

	str, err = clipboard.ReadString()
	assert.NoErr(t, err)
	assert.NotEmpty(t, str)
	assert.Eq(t, src, str)
}

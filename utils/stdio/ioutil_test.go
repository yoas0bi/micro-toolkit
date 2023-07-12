package stdio_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/fsutil"
	"github.com/yoas0bi/micro-toolkit/utils/stdio"
	"github.com/yoas0bi/micro-toolkit/utils/testutil"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestQuietFprint(t *testing.T) {
	buf := testutil.NewBuffer()

	stdio.QuietFprint(buf, "hi, inhere")
	assert.Eq(t, "hi, inhere", buf.ResetAndGet())

	stdio.QuietFprintf(buf, "hi, %s", "inhere")
	assert.Eq(t, "hi, inhere", buf.ResetAndGet())

	stdio.QuietFprintln(buf, "hi, inhere")
	assert.Eq(t, "hi, inhere\n", buf.ResetAndGet())
}

func TestQuietWriteString(t *testing.T) {
	buf := new(bytes.Buffer)
	stdio.QuietWriteString(buf, "inhere")

	assert.Eq(t, "inhere", buf.String())
}

func TestDiscardReader(t *testing.T) {
	sr := strings.NewReader("hello")
	fsutil.DiscardReader(sr)

	assert.Empty(t, stdio.MustReadReader(sr))
}

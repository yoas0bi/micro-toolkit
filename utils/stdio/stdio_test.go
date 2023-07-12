package stdio_test

import (
	"strings"
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/stdio"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestNewScanner(t *testing.T) {
	s := stdio.NewScanner("hi\nmy\nname\nis\ninhere")

	var ss []string
	// scan each line
	for s.Scan() {
		ss = append(ss, s.Text())
	}

	assert.Len(t, ss, 5)
	assert.Eq(t, "hi", ss[0])
}

func TestNewIOReader(t *testing.T) {
	assert.Panics(t, func() {
		stdio.NewIOReader([]int{23})
	})

	r := stdio.NewIOReader("hi")
	assert.Eq(t, "hi", stdio.ReadString(r))
	r = stdio.NewIOReader([]byte("hi"))
	assert.Eq(t, "hi", stdio.ReadString(r))
	r = stdio.NewIOReader(strings.NewReader("hi"))
	assert.Eq(t, "hi", stdio.ReadString(r))
}

func TestWriteBytes(t *testing.T) {
	stdio.WriteBytes([]byte("hi,"))
	stdio.WriteString("inhere.")
	stdio.Writeln("welcome")
}

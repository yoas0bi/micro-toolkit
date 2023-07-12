package cmdr_test

import (
	"fmt"
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/sysutil/cmdr"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestNewCmd(t *testing.T) {
	c := cmdr.NewCmd("ls").
		WithArg("-l").
		WithArgs([]string{"-h"}).
		AddArg("-a").
		AddArgf("%s", "./")

	assert.Eq(t, "ls", c.BinName())
	assert.Eq(t, "ls", c.IDString())
	assert.StrContains(t, "ls", c.BinOrPath())
	assert.NotContains(t, c.OnlyArgs(), "ls")

	c.OnBefore(func(c *cmdr.Cmd) {
		assert.Eq(t, "ls -l -h -a ./", c.Cmdline())
	})

	out := c.SafeOutput()
	fmt.Println(out)
	assert.NotEmpty(t, out)
	assert.NotEmpty(t, cmdr.OutputLines(out))
	assert.NotEmpty(t, cmdr.FirstLine(out))

	c.ResetArgs()
	assert.Len(t, c.Args, 1)
	assert.Empty(t, c.OnlyArgs())
}

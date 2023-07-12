package testutil_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/testutil"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestDiscardStdout(t *testing.T) {
	err := testutil.DiscardStdout()
	fmt.Println("Hello, playground")
	str := testutil.RestoreStdout()

	assert.NoErr(t, err)
	assert.Eq(t, "", str)
}

func TestRewriteStdout(t *testing.T) {
	testutil.RewriteStdout()

	assert.Eq(t, "", testutil.RestoreStdout())

	testutil.RewriteStdout()
	fmt.Println("Hello, playground")
	msg := testutil.RestoreStdout()

	assert.Eq(t, "Hello, playground\n", msg)
}

func TestRewriteStderr(t *testing.T) {
	testutil.RewriteStderr()
	assert.Eq(t, "", testutil.RestoreStderr())

	testutil.RewriteStderr()
	_, err := fmt.Fprint(os.Stderr, "Hello, playground")
	msg := testutil.RestoreStderr()

	assert.NoErr(t, err)
	assert.Eq(t, "Hello, playground", msg)
}

func TestNewDirEnt(t *testing.T) {
	de := testutil.NewDirEnt("testdata/some.txt")
	assert.NotEmpty(t, de)
	assert.False(t, de.IsDir())
}

package fsutil_test

import (
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/fsutil"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestMustCopyFile(t *testing.T) {
	srcPath := "./testdata/cp-file-src.txt"
	dstPath := "./testdata/cp-file-dst.txt"

	assert.NoErr(t, fsutil.RmIfExist(srcPath))
	assert.NoErr(t, fsutil.RmFileIfExist(dstPath))

	_, err := fsutil.PutContents(srcPath, "hello")
	assert.NoErr(t, err)

	fsutil.MustCopyFile(srcPath, dstPath)
	assert.Eq(t, []byte("hello"), fsutil.GetContents(dstPath))
	assert.Eq(t, "hello", fsutil.ReadString(dstPath))

	str, err := fsutil.ReadStringOrErr(dstPath)
	assert.NoErr(t, err)
	assert.Eq(t, "hello", str)

	assert.NoErr(t, fsutil.RmFileIfExist(srcPath))
	assert.NoErr(t, fsutil.RmIfExist(srcPath)) // repeat delete
}

func TestWriteFile(t *testing.T) {
	tempFile := "./testdata/write-file.txt"

	err := fsutil.WriteFile(tempFile, []byte("hello\ngolang"), 0644)
	assert.NoErr(t, err)
	assert.Eq(t, []byte("hello\ngolang"), fsutil.MustReadFile(tempFile))

	// write invalid data
	assert.Panics(t, func() {
		_ = fsutil.WriteFile(tempFile, []string{"hello"}, 0644)
	})
}

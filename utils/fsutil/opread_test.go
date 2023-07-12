package fsutil_test

import (
	"strings"
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/fsutil"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestDiscardReader(t *testing.T) {
	sr := strings.NewReader("hello")
	bs, err := fsutil.ReadOrErr(sr)
	assert.NoErr(t, err)
	assert.Eq(t, []byte("hello"), bs)

	sr = strings.NewReader("hello")
	assert.Eq(t, []byte("hello"), fsutil.GetContents(sr))

	sr = strings.NewReader("hello")
	fsutil.DiscardReader(sr)

	assert.Empty(t, fsutil.ReadReader(sr))
	assert.Empty(t, fsutil.ReadAll(sr))

}

func TestGetContents(t *testing.T) {
	fpath := "./testdata/get-contents.txt"
	assert.NoErr(t, fsutil.RmFileIfExist(fpath))

	_, err := fsutil.PutContents(fpath, "hello")
	assert.NoErr(t, err)

	assert.Nil(t, fsutil.ReadExistFile("/path-not-exist"))
	assert.Eq(t, []byte("hello"), fsutil.ReadExistFile(fpath))

	assert.Panics(t, func() {
		fsutil.GetContents(45)
	})
	assert.Panics(t, func() {
		fsutil.ReadFile("/path-not-exist")
	})
}

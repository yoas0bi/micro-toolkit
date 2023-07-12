package fsutil_test

import (
	"bytes"
	"io/fs"
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/basefn"
	"github.com/yoas0bi/micro-toolkit/utils/fsutil"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestMain(m *testing.M) {
	err := fsutil.RemoveSub("./testdata", func(fPath string, ent fs.DirEntry) bool {
		return fsutil.PathMatch(ent.Name(), "*.txt")
	})
	basefn.MustOK(err)

	m.Run()
}

func TestMimeType(t *testing.T) {
	assert.Eq(t, "", fsutil.DetectMime(""))
	assert.Eq(t, "", fsutil.MimeType("not-exist"))
	assert.Eq(t, "image/jpeg", fsutil.MimeType("testdata/test.jpg"))

	buf := new(bytes.Buffer)
	buf.Write([]byte("\xFF\xD8\xFF"))
	assert.Eq(t, "image/jpeg", fsutil.ReaderMimeType(buf))
	buf.Reset()

	buf.Write([]byte("text"))
	assert.Eq(t, "text/plain; charset=utf-8", fsutil.ReaderMimeType(buf))
	buf.Reset()

	buf.Write([]byte(""))
	assert.Eq(t, "", fsutil.ReaderMimeType(buf))
	buf.Reset()

	assert.True(t, fsutil.IsImageFile("testdata/test.jpg"))
	assert.False(t, fsutil.IsImageFile("testdata/not-exists"))
}

func TestTempDir(t *testing.T) {
	dir, err := fsutil.TempDir("testdata", "temp.*")
	assert.NoErr(t, err)
	assert.True(t, fsutil.IsDir(dir))
	assert.NoErr(t, fsutil.Remove(dir))
}

func TestSplitPath(t *testing.T) {
	dir, file := fsutil.SplitPath("/path/to/dir/some.txt")
	assert.Eq(t, "/path/to/dir/", dir)
	assert.Eq(t, "some.txt", file)
}

func TestToAbsPath(t *testing.T) {
	assert.Eq(t, "", fsutil.ToAbsPath(""))
	assert.Eq(t, "/path/to/dir/", fsutil.ToAbsPath("/path/to/dir/"))
	assert.Neq(t, "~/path/to/dir", fsutil.ToAbsPath("~/path/to/dir"))
	assert.Neq(t, ".", fsutil.ToAbsPath("."))
	assert.Neq(t, "..", fsutil.ToAbsPath(".."))
	assert.Neq(t, "./", fsutil.ToAbsPath("./"))
	assert.Neq(t, "../", fsutil.ToAbsPath("../"))
}

func TestSlashPath(t *testing.T) {
	assert.Eq(t, "/path/to/dir", fsutil.SlashPath("/path/to/dir"))
	assert.Eq(t, "/path/to/dir", fsutil.UnixPath("/path/to/dir"))
	assert.Eq(t, "/path/to/dir", fsutil.UnixPath("\\path\\to\\dir"))
}

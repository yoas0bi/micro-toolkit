package fsutil

import (
	"io"
	"os"

	"github.com/yoas0bi/micro-toolkit/utils/basefn"
)

// OSTempFile create a temp file on os.TempDir()
//
// Usage:
//
//	fsutil.OSTempFile("example.*.txt")
func OSTempFile(pattern string) (*os.File, error) {
	return os.CreateTemp(os.TempDir(), pattern)
}

// TempFile is like os.CreateTemp, but can custom temp dir.
//
// Usage:
//
//	fsutil.TempFile("", "example.*.txt")
func TempFile(dir, pattern string) (*os.File, error) {
	return os.CreateTemp(dir, pattern)
}

// OSTempDir creates a new temp dir on os.TempDir and return the temp dir path
//
// Usage:
//
//	fsutil.OSTempDir("example.*")
func OSTempDir(pattern string) (string, error) {
	return os.MkdirTemp(os.TempDir(), pattern)
}

// TempDir creates a new temp dir and return the temp dir path
//
// Usage:
//
//	fsutil.TempDir("", "example.*")
//	fsutil.TempDir("testdata", "example.*")
func TempDir(dir, pattern string) (string, error) {
	return os.MkdirTemp(dir, pattern)
}

// ************************************************************
//	write, copy files
// ************************************************************

// PutContents create file and write contents to file at once.
//
// data type allow: string, []byte, io.Reader
//
// Tip: file flag default is FsCWTFlags (override write)
//
// Usage:
//
//	fsutil.PutContents(filePath, contents, fsutil.FsCWAFlags) // append write
func PutContents(filePath string, data any, fileFlag ...int) (int, error) {
	f, err := QuickOpenFile(filePath, basefn.FirstOr(fileFlag, FsCWTFlags))
	if err != nil {
		return 0, err
	}

	return WriteOSFile(f, data)
}

// WriteFile create file and write contents to file, can set perm for file.
//
// data type allow: string, []byte, io.Reader
//
// Tip: file flag default is FsCWTFlags (override write)
//
// Usage:
//
//	fsutil.WriteFile(filePath, contents, fsutil.DefaultFilePerm, fsutil.FsCWAFlags)
func WriteFile(filePath string, data any, perm os.FileMode, fileFlag ...int) error {
	flag := basefn.FirstOr(fileFlag, FsCWTFlags)
	f, err := OpenFile(filePath, flag, perm)
	if err != nil {
		return err
	}

	_, err = WriteOSFile(f, data)
	return err
}

// WriteOSFile write data to give os.File, then close file.
//
// data type allow: string, []byte, io.Reader
func WriteOSFile(f *os.File, data any) (n int, err error) {
	switch typData := data.(type) {
	case []byte:
		n, err = f.Write(typData)
	case string:
		n, err = f.WriteString(typData)
	case io.Reader: // eg: buffer
		var n64 int64
		n64, err = io.Copy(f, typData)
		n = int(n64)
	default:
		_ = f.Close()
		panic("WriteFile: data type only allow: []byte, string, io.Reader")
	}

	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return n, err
}

// CopyFile copy a file to another file path.
func CopyFile(srcPath, dstPath string) error {
	srcFile, err := os.OpenFile(srcPath, FsRFlags, 0)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// create and open file
	dstFile, err := QuickOpenFile(dstPath, FsCWTFlags)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// MustCopyFile copy file to another path.
func MustCopyFile(srcPath, dstPath string) {
	err := CopyFile(srcPath, dstPath)
	if err != nil {
		panic(err)
	}
}

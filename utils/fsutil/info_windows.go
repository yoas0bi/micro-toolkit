package fsutil

import (
	"path/filepath"

	"github.com/yoas0bi/micro-toolkit/utils/internal/comfunc"
)

// Realpath returns the shortest path name equivalent to path by purely lexical processing.
func Realpath(pathStr string) string {
	pathStr = comfunc.ExpandHome(pathStr)

	if !IsAbsPath(pathStr) {
		pathStr = JoinSubPaths(comfunc.Workdir(), pathStr)
	}
	return filepath.Clean(pathStr)
}

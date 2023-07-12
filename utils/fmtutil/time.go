package fmtutil

import (
	"github.com/yoas0bi/micro-toolkit/utils/basefn"
)

// HowLongAgo format a seconds, get how lang ago
func HowLongAgo(sec int64) string {
	return basefn.HowLongAgo(sec)
}

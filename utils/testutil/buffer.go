package testutil

import (
	"github.com/yoas0bi/micro-toolkit/utils/byteutil"
)

// Buffer wrap and extends the bytes.Buffer
type Buffer = byteutil.Buffer

// NewBuffer instance
func NewBuffer() *byteutil.Buffer {
	return byteutil.NewBuffer()
}

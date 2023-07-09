package defaultvalue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yoas0bi/micro-toolkit/tools/defaultvalue"
)

func TestWithDefault(t *testing.T) {
	should := assert.New(t)

	v := defaultvalue.WithDefault("", "default")
	should.Equal("default", v)
	v = defaultvalue.WithDefault("has data", "default")
	should.Equal("has data", v)
}

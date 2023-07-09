package enum_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yoas0bi/micro-toolkit/cmd/mcube/enum"
)

func TestGenerate(t *testing.T) {
	should := assert.New(t)
	code, err := enum.G.Generate("../../../examples/enum/enum.go")
	t.Log(string(code))
	should.NoError(err)
}

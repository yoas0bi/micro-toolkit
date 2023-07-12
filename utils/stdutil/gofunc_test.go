package stdutil_test

import (
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/dump"
	"github.com/yoas0bi/micro-toolkit/utils/stdutil"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestFuncName(t *testing.T) {
	name := stdutil.FuncName(stdutil.PkgName)
	assert.Eq(t, "github.com/yoas0bi/micro-toolkit/utils/stdutil.PkgName", name)

	name = stdutil.FuncName(stdutil.PanicIfErr)
	assert.Eq(t, "github.com/yoas0bi/micro-toolkit/utils/stdutil.PanicIfErr", name)
}

func TestPkgName(t *testing.T) {
	name := stdutil.PkgName(stdutil.FuncName(stdutil.PanicIfErr))
	assert.Eq(t, "github.com/yoas0bi/micro-toolkit/utils/stdutil", name)
}

func TestFullFcName_Parse(t *testing.T) {
	fullName := stdutil.FuncName(stdutil.PanicIfErr)

	ffn := stdutil.FullFcName{FullName: fullName}
	ffn.Parse()
	assert.Eq(t, fullName, ffn.String())
	assert.Eq(t, "stdutil", ffn.PkgName())
	assert.Eq(t, "PanicIfErr", ffn.FuncName())
	assert.Eq(t, "github.com/yoas0bi/micro-toolkit/utils/stdutil", ffn.PkgPath())
	dump.P(ffn)

	st := stdutil.FullFcName{}
	fullName = stdutil.FuncName(st.FuncName)

	ffn = stdutil.FullFcName{FullName: fullName}
	ffn.Parse()
	assert.Eq(t, "(*FullFcName).FuncName-fm", ffn.FuncName())
	dump.P(ffn)
}

func TestCutFuncName(t *testing.T) {
	fullName := stdutil.FuncName(stdutil.PanicIfErr)

	pkgPath, funcName := stdutil.CutFuncName(fullName)
	assert.Eq(t, "PanicIfErr", funcName)
	assert.Eq(t, "github.com/yoas0bi/micro-toolkit/utils/stdutil", pkgPath)
}

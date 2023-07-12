package uuid

import (
	"github.com/yoas0bi/micro-toolkit/utils/helper"
	"testing"
)

func TestUuid(t *testing.T) {
	uuid := helper.TUuid.Uuid()
	if uuid == "false" {
		t.Error("uuid unit test fail")
	}
}

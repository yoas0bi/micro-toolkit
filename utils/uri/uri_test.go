package uri

import (
	"github.com/yoas0bi/micro-toolkit/utils/helper"
	"testing"
)

func TestParseUriQueryToMap(t *testing.T) {

	m := helper.TUri.ParseUriQueryToMap(helper.exampleUriStr)
	if _, ok := m["Av"]; !ok {
		t.Errorf("ParseUriQueryToMap unit test fail")
	}
}

func TestGetQueryParams(t *testing.T) {
	m := make(UriValues)
	m, mError := helper.TUri.ParseUriQuery(m, helper.exampleUriStr)
	if mError != nil {
		t.Errorf("parseQuery Errors:%v \n", mError)
	}
	av := helper.TUri.GetQueryParams(m, "av")
	if av == "" {
		t.Errorf("the values of %v is not %v \n", "5.3.5", "")
	}
}

func TestGetDomain(t *testing.T) {
	actual := ""
	for _, test := range helper.exampleUriTests {
		actual = helper.TUri.GetDomain(test.param, test.isMain)
		if actual != test.expected {
			t.Errorf("Expected GetDomain(%q) to be %v, got is %v \n", test.param, test.expected, actual)
		}
	}
	helper.TUri.GetDomain("123456")
}

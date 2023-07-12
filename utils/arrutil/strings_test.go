package arrutil_test

import (
	"fmt"
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/arrutil"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestStringsToSlice(t *testing.T) {
	is := assert.New(t)

	as := arrutil.StringsToSlice([]string{"1", "2"})
	is.Eq(`[]interface {}{"1", "2"}`, fmt.Sprintf("%#v", as))
}

func TestStringsToInts(t *testing.T) {
	is := assert.New(t)

	ints, err := arrutil.StringsToInts([]string{"1", "2"})
	is.Nil(err)
	is.Eq("[]int{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = arrutil.StringsToInts([]string{"a", "b"})
	is.Err(err)

	ints = arrutil.StringsAsInts([]string{"1", "2"})
	is.Eq("[]int{1, 2}", fmt.Sprintf("%#v", ints))
	is.Nil(arrutil.StringsAsInts([]string{"abc"}))
}

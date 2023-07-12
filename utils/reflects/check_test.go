package reflects_test

import (
	"reflect"
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/reflects"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestIsNil(t *testing.T) {
	assert.False(t, reflects.IsNil(reflect.ValueOf(nil)))

	var v *reflects.Value
	assert.True(t, reflects.IsNil(reflect.ValueOf(v)))
}

func TestIsFunc(t *testing.T) {
	assert.False(t, reflects.IsFunc(nil))
	assert.True(t, reflects.IsFunc(reflects.HasChild))
}

func TestHasChild(t *testing.T) {
	is := assert.New(t)

	is.True(reflects.HasChild(reflect.ValueOf([]int{23})))
	is.False(reflects.HasChild(reflect.ValueOf("abc")))
}

func TestIsEqual(t *testing.T) {
	is := assert.New(t)

	is.False(reflects.IsEqual(nil, "abc"))
	is.False(reflects.IsEqual("abc", nil))
	is.False(reflects.IsEqual("abc", 123))
	is.True(reflects.IsEqual(123, 123))
	is.True(reflects.IsEqual([]byte{}, []byte{}))
	is.True(reflects.IsEqual([]byte("abc"), []byte("abc")))
	is.False(reflects.IsEqual([]byte("abc"), 123))
}

func TestIsEmpty(t *testing.T) {
	is := assert.New(t)

	is.True(reflects.IsEmpty(reflect.ValueOf(nil)))
	is.True(reflects.IsEmpty(reflect.ValueOf(false)))
	is.True(reflects.IsEmpty(reflect.ValueOf("")))
	is.True(reflects.IsEmpty(reflect.ValueOf([]string{})))
	is.True(reflects.IsEmpty(reflect.ValueOf(map[int]string{})))
	is.True(reflects.IsEmpty(reflect.ValueOf(0)))
	is.True(reflects.IsEmpty(reflect.ValueOf(uint(0))))
	is.True(reflects.IsEmpty(reflect.ValueOf(float32(0))))

	type T struct {
		v any //lint:ignore U1000 for test
	}
	rv := reflect.ValueOf(T{}).Field(0)
	is.True(reflects.IsEmpty(rv))
}

func TestIsEmptyValue(t *testing.T) {
	is := assert.New(t)

	is.True(reflects.IsEmptyValue(reflect.ValueOf(nil)))
	is.True(reflects.IsEmptyValue(reflect.ValueOf("")))
	is.True(reflects.IsEmptyValue(reflect.ValueOf(false)))
	is.True(reflects.IsEmptyValue(reflect.ValueOf([]string{})))
	is.True(reflects.IsEmptyValue(reflect.ValueOf(map[int]string{})))
	is.True(reflects.IsEmptyValue(reflect.ValueOf(0)))
	is.True(reflects.IsEmptyValue(reflect.ValueOf(uint(0))))
	is.True(reflects.IsEmptyValue(reflect.ValueOf(float32(0))))

	type T struct {
		v any //lint:ignore U1000 for test
	}
	rv := reflect.ValueOf(T{}).Field(0)
	is.True(reflects.IsEmptyValue(rv))

	rv = reflect.ValueOf(&T{v: "abc"})
	is.False(reflects.IsEmptyValue(rv))
}

func TestIsSimpleKind(t *testing.T) {
	testCases := []struct {
		name     string
		input    reflect.Kind
		expected bool
	}{
		{"invalid kind", reflect.Invalid, false},
		{"string kind", reflect.String, true},
		{"float64 kind", reflect.Float64, true},
		{"bool kind", reflect.Bool, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if reflects.IsSimpleKind(tc.input) != tc.expected {
				t.Errorf("expected %v but got %v", tc.expected, !tc.expected)
			}
		})
	}
}

func TestIsAnyInt(t *testing.T) {
	// test for IsAnyInt
	assert.True(t, reflects.IsAnyInt(reflect.Int))
	assert.True(t, reflects.IsAnyInt(reflect.Int8))
	assert.True(t, reflects.IsAnyInt(reflect.Uint))
	assert.False(t, reflects.IsAnyInt(reflect.Func))

	// test for IsIntx
	assert.True(t, reflects.IsIntx(reflect.Int))
	assert.False(t, reflects.IsIntx(reflect.Uint))

	// test for IsUintX()
	assert.True(t, reflects.IsUintX(reflect.Uint16))
	assert.False(t, reflects.IsUintX(reflect.Int))
}

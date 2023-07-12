package int

import (
	"github.com/yoas0bi/micro-toolkit/utils/helper"
	"testing"
)

func TestAbs(t *testing.T) {
	got1 := -1
	want1 := int64(1)
	abs1 := helper.TInt.Abs(int64(got1))
	if abs1 != want1 {
		t.Errorf("The Abs values of %v is not %v\n", abs1, want1)
	}

	got2 := 0
	want2 := int64(0)
	abs2 := helper.TInt.Abs(int64(got2))
	if abs2 != want2 {
		t.Errorf("The Abs values of %v is not %v\n", abs2, want2)
	}

	got3 := 1
	want3 := int64(1)
	abs3 := helper.TInt.Abs(int64(got3))
	if abs3 != want3 {
		t.Errorf("The Abs values of %v is not %v\n", abs3, want3)
	}
}

func TestIsOddIsEven(t *testing.T) {
	res1 := helper.TValidate.IsOdd(-1)
	res2 := helper.TValidate.IsOdd(0)
	res3 := helper.TValidate.IsEven(2)
	res4 := helper.TValidate.IsEven(-3)

	if !res1 || res2 {
		t.Error("IsOdd unit test fail")
		return
	} else if !res3 || res4 {
		t.Error("IsEven unit test fail")
		return
	}
}

func TestInRangeInt(t *testing.T) {
	for _, test := range helper.exampleIntRange {
		actual := helper.TValidate.IsRangeInt(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRangeInt(%v, %v, %v) to be %v, got %v using type int", test.param, test.left, test.right, test.expected, actual)
		}
	}
}

func TestAverageInt(t *testing.T) {
	var res1, res2, res3 float64

	res1 = helper.TInt.AverageInt()
	res2 = helper.TInt.AverageInt(1)
	res3 = helper.TInt.AverageInt(1, 2, 3, 4, 5, 6)

	if res1 != 0 || int(res2) != 1 || helper.TInt.NumberFormat(res3, 2, ".", "") != "3.50" {
		t.Error("AverageInt unit test fail")
		return
	}
}

func TestSumInt(t *testing.T) {
	sum := helper.TInt.SumInt(0, 1, -2, 3, 5)
	if sum != 7 {
		t.Error("SumInt unit test fail")
		return
	}
}

func TestMinInt(t *testing.T) {
	nums := []int{0, 3, -4, 5, 9}
	min := helper.TInt.MinInt(nums...)
	if min != -4 {
		t.Error("MinInt unit test fail")
		return
	}

	min = helper.TInt.MinInt(-1)
	if min != -1 {
		t.Error("MinInt unit test fail")
		return
	}
}

func TestMaxInt(t *testing.T) {
	nums := []int{-4, 0, 3, 9}
	max := helper.TInt.MaxInt(nums...)
	if max != 9 {
		t.Error("MaxInt unit test fail")
		return
	}

	max = helper.TInt.MaxInt(-1)
	if max != -1 {
		t.Error("MaxInt unit test fail")
		return
	}
}

func TestNumberFormat(t *testing.T) {
	num1 := 123.4567890
	num2 := -123.4567890
	num3 := 123456789.1234567890
	res1 := helper.TInt.NumberFormat(num1, 3, ".", "")
	res2 := helper.TInt.NumberFormat(num2, 0, ".", "")
	res3 := helper.TInt.NumberFormat(num3, 6, ".", ",")
	if res1 != "123.457" || res2 != "-123" || res3 != "123,456,789.123457" {
		t.Error("NumberFormat unit test fail")
		return
	}
}

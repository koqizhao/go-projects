package mathext

import "testing"
import "math"

const _TEST_INCREMENT = 0.1
const _TEST_DOUBLE_INCREMENT = 0.2
const _TEST_ITERATION_TIMES = 1000
const _IGNORE_INT_VALUE = int(math.MinInt32)

var _float64Values = []float64{1, 1.5, 1.1, 1.0, 1.7, 2.0, 0, -1, -1.1, -1.5, -1.7, -2.0}
var _roundedInt64Values = []int64{1, 2, 1, 1, 2, 2, 0, -1, -1, -2, -2, -2}
var _float32Values = []float32{1, 1.5, 1.1, 1.0, 1.7, 2.0, 0, -1, -1.1, -1.5, -1.7, -2.0}
var _roundedInt32Values = []int32{1, 2, 1, 1, 2, 2, 0, -1, -1, -2, -2, -2}

func TestSum(t *testing.T) {
	r := Sum(_float64Values...)
	if !AreEqualFloat(1, r) {
		t.Error("Expected", 1, "Actual", r)
	}
}

func TestInt64Overflow(t *testing.T) {
	var y float64 = math.MaxInt64 + 1
	if int64(y) != math.MinInt64 {
		t.Error("MaxInt64 overflow", int64(y))
	}

	y = math.MinInt64 - 1
	if int64(y) != math.MinInt64 {
		t.Error("MinInt64 overflow", int64(y))
	}
}

func TestAreEqualFloat(t *testing.T) {
	testAreEqualFloat(t, _IGNORE_INT_VALUE)
}

func TestAreEqualFloatWithAccuracy(t *testing.T) {
	accuracyList := []int{-1, 0, _DEFAULT_FLOAT_ACCURACY - 1, _DEFAULT_FLOAT_ACCURACY}
	for _, v := range accuracyList {
		testAreEqualFloat(t, v)
	}
}

func TestRound(t *testing.T) {
	for i, l := 0, len(_float64Values); i < l; i++ {
		if Round(_float64Values[i]) != _roundedInt64Values[i] {
			t.Error("Float", _float64Values[i], "Int", _roundedInt64Values[i], "Expected", "Equal")
			return
		}
	}
}

func TestRoundFloat32(t *testing.T) {
	for i, l := 0, len(_float32Values); i < l; i++ {
		if RoundFloat32(_float32Values[i]) != _roundedInt32Values[i] {
			t.Error("Float", _float32Values[i], "Int", _roundedInt32Values[i], "Expected", "Equal")
			return
		}
	}
}

func testAreEqualFloat(t *testing.T, accuracy int) {
	var areEqual bool
	for x, y, i := 0.0, 0.0, 0; i < _TEST_ITERATION_TIMES; i++ {
		x += _TEST_INCREMENT
		if i%2 == 0 {
			y += _TEST_DOUBLE_INCREMENT
			if accuracy == _IGNORE_INT_VALUE {
				areEqual = AreEqualFloat(x, y)
			} else {
				areEqual = AreEqualFloatWithAccuracy(x, y, accuracy)
			}
			if areEqual {
				t.Error("x", x, "y", y, "Expected", "NotEqual")
				return
			}
		} else {
			if accuracy == _IGNORE_INT_VALUE {
				areEqual = AreEqualFloat(x, y)
			} else {
				areEqual = AreEqualFloatWithAccuracy(x, y, accuracy)
			}
			if !areEqual {
				t.Error("x", x, "y", y, "Expected", "Equal")
				return
			}
		}
	}
}

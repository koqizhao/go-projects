package mathext

import "testing"
import "math"

const _TEST_INCREMENT = 0.1
const _TEST_DOUBLE_INCREMENT = 0.2
const _TEST_ITERATION_TIMES = 1000
const _IGNORE_INT_VALUE = int(math.MinInt32)

func TestAreEqualFloat(t *testing.T) {
    testAreEqualFloat(t, _IGNORE_INT_VALUE)
}

func TestAreEqualFloatWithAccuracy(t *testing.T) {
    accuracyList := []int { -1, 0, _DEFAULT_FLOAT_ACCURACY -1, _DEFAULT_FLOAT_ACCURACY }
    for _, v := range accuracyList {
        testAreEqualFloat(t, v)
    }
}

func testAreEqualFloat(t *testing.T, accuracy int) {
    var areEqual bool
    for x, y, i := 0.0, 0.0, 0; i < _TEST_ITERATION_TIMES; i++ {
        x += _TEST_INCREMENT
        if i % 2 == 0 {
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

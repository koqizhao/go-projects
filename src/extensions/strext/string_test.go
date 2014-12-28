package strext

import "testing"
import "unicode/utf8"

var _number = 123
var _s = ToString(123)
var _width = 10
var _char = '@'
var _padding = Mul(string(_char), _width - utf8.RuneCountInString(_s))
var _expectedLeftPadding = _padding + _s
var _expectedRightPadding = _s + _padding

func TestPadLeft(t *testing.T) {
    if r := PadLeft(_number, _width, _char); r != _expectedLeftPadding {
        t.Error("Actual", r, "Expected", _expectedLeftPadding)
    }
}

func TestPadRight(t *testing.T) {
    if r := PadRight(_number, _width, _char); r != _expectedRightPadding {
        t.Error("Actual", r, "Expected", _expectedRightPadding)
    }
}

func TestSimplifyWhitespace(t *testing.T) {
    original, expected := " x x   y y   xxxy\ty\t  ", "x x y y xxxy y"
    if actual := SimplifyWhitespace(original); actual != expected {
        t.Error("Actual", actual, "Expected", expected, "Original", original)
    }
}

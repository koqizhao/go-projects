package mathext

import "math"

const DEFAULT_FLOAT_ACCURACY int = 8
const _DEFAULT_FLOAT_ACCURACY int = 8
const _MAX_FLOAT_ACCURACY int = 128

func AreEqualFloat(x, y float64) bool {
    return AreEqualFloatWithAccuracy(x, y, _DEFAULT_FLOAT_ACCURACY)
}

func AreEqualFloatWithAccuracy(x, y float64, accuracy int) bool {
    return math.Abs(x - y) < calculateMarginOfError(accuracy)
}

func calculateMarginOfError(accuracy int) float64 {
    if accuracy <= 0 {
        accuracy = _DEFAULT_FLOAT_ACCURACY
    } else if accuracy > _MAX_FLOAT_ACCURACY {
        accuracy = _MAX_FLOAT_ACCURACY
    }

    result := 1.0
    for i := 0; i < accuracy; i++ {
        result /= 10
    }
    return result
}

package mathext

import "math"
import "fmt"

const DEFAULT_FLOAT_ACCURACY int = 8
const _DEFAULT_FLOAT_ACCURACY int = 8
const _MAX_FLOAT_ACCURACY int = 128

func AreEqualFloat(x, y float64) bool {
    return AreEqualFloatWithAccuracy(x, y, _DEFAULT_FLOAT_ACCURACY)
}

func AreEqualFloatWithAccuracy(x, y float64, accuracy int) bool {
    return math.Abs(x - y) < calculateMarginOfError(accuracy)
}

func Round(x float64) int64 {
    whole, fraction := math.Modf(x)
    if whole < math.MinInt64 || whole > math.MaxInt64 {
        panic(fmt.Sprintf("%g is out of the int64 range", x))
    }
    if math.Abs(fraction) >= 0.5 {
        if x >= 0 {
            whole++
        } else {
            whole--
        }
    }
    return int64(whole)
}

func RoundFloat32(x float32) int32 {
    return int32(Round(float64(x)))
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

package mathext

import "math"
import "fmt"

const _DEFAULT_FLOAT_ACCURACY int = 8
const _MAX_FLOAT_ACCURACY int = 128

func AreEqualFloat(x, y float64) bool {
	return AreEqualFloatWithAccuracy(x, y, _DEFAULT_FLOAT_ACCURACY)
}

func AreEqualFloatWithAccuracy(x, y float64, accuracy int) bool {
	return math.Abs(x-y) < calculateMarginOfError(accuracy)
}

func Round(x float64) int64 {
	whole, fraction := math.Modf(x)
	if whole < math.MinInt64 || whole > math.MaxInt64 {
		panic(fmt.Sprintf("%g is out of the int64 range", x))
	}
	if math.Abs(fraction) >= 0.5 {
		if whole == math.MinInt64 || whole == math.MaxInt64 {
			panic(fmt.Sprintf("%g is out of the int64 range", x))
		}
		if x >= 0 {
			whole++
		} else {
			whole--
		}
	}
	return int64(whole)
}

func RoundFloat32(x float32) int32 {
	rounded := int32(Round(float64(x)))
	if rounded > math.MaxInt32 || rounded < math.MinInt32 {
		panic(fmt.Sprintf("%g is out of the int32 range", x))
	}
	return int32(rounded)
}

func Sum(arr []float64) float64 {
	result := 0.0
	for i, l := 0, len(arr); i < l; i++ {
		result += arr[i]
	}
	return result
}

func Avg(arr []float64) float64 {
	if l := len(arr); l == 0 {
		return 0
	} else {
		return Sum(arr) / float64(l)
	}
}

func Median(arr []float64) float64 {
	l := len(arr)
	if l == 0 {
		return 0
	}

	result := arr[l/2]
	if l%2 == 0 {
		result = (result + arr[l/2-1]) / 2
	}
	return result
}

func calculateMarginOfError(accuracy int) float64 {
	if accuracy <= 0 {
		accuracy = _DEFAULT_FLOAT_ACCURACY
	} else if accuracy > _MAX_FLOAT_ACCURACY {
		accuracy = _MAX_FLOAT_ACCURACY
	}

	return math.Pow10(-accuracy)
}

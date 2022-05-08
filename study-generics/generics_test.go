package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_sum(t *testing.T) {
	result := sum(1, 2, 3)
	fmt.Printf("result: %d\n", result)
}

func Test_sumNumber(t *testing.T) {
	result := sumNumber(1.1, 2.2, 3.3)
	fmt.Printf("result: %v\n", result)
}

func Test_comparable(t *testing.T) {
	var v any
	fmt.Printf("result: %v\n", v)
}

func Test_genericTypes(t *testing.T) {
	var v1, v2 GenericStruct[int]
	v1.V1, v1.V2, v2.V1, v2.V2 = 1, 2, 3, 4
	v3 := v1.sum2(v1, v2)
	var v4 GenericInterface[int] = v3
	v5 := v4.sum2(v1, v3)
	fmt.Printf("v1: %v, v2: %v, v3: %v, v4: %v, v5: %v\n", v1, v2, v3, v4, v5)

	fmt.Printf("type: %v\n", reflect.TypeOf(v5))
}

func Test_any(t *testing.T) {
	var x any
	fmt.Printf("any: %v\n", x)
}

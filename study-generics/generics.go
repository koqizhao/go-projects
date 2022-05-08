package main

func sum[V int | float32](values ...V) V {
	var result V
	for _, v := range values {
		result += v
	}

	return result
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func sumNumber[V Number](values ...V) V {
	var result V
	for _, v := range values {
		result += v
	}

	return result
}

type GenericInterface[T Number] interface {
	sum(v1 T, v2 T) T
	sum2(v1 GenericStruct[T], v2 GenericStruct[T]) GenericStruct[T]
}

type GenericStruct[T Number] struct {
	V1 T
	V2 T
}

func (v GenericStruct[T]) sum(v1 T, v2 T) T {
	return v1 + v2
}

func (v GenericStruct[T]) sum2(v1 GenericStruct[T], v2 GenericStruct[T]) GenericStruct[T] {
	return GenericStruct[T]{
		V1: v1.V1 + v2.V1,
		V2: v1.V2 + v2.V2,
	}
}

/*
func (v GenericStruct[T]) sum3[T1 int | float64](v1 T1, v2 T1) T1 {
	return v1 + v2
}
*/

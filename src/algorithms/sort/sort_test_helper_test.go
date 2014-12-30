package sort

import "testing"

var originalData = [][]int{
	[]int{1, 2, 3, 4, 5},
	[]int{5, 4, 3, 2, 1},
	[]int{1, 1, 1, 1, 1},
	[]int{1, 2, 3, 2, 1},
	[]int{1},
	[]int{},
	[]int{1, 1, 1, 1, 2},
	[]int{2, 1, 1, 1, 1},
	[]int{0, 0, 0, 0, 0},
}

var expectedData = [][]int{
	[]int{1, 2, 3, 4, 5},
	[]int{1, 2, 3, 4, 5},
	[]int{1, 1, 1, 1, 1},
	[]int{1, 1, 2, 2, 3},
	[]int{1},
	[]int{},
	[]int{1, 1, 1, 1, 2},
	[]int{1, 1, 1, 1, 2},
	[]int{0, 0, 0, 0, 0},
}

type sortMethod func(arr *[]int)

func areEqual(arr1, arr2 *[]int) bool {
	if arr1 == nil && arr2 == nil {
		return true
	}

	if arr1 == nil || arr2 == nil {
		return false
	}

	if len(*arr1) != len(*arr2) {
		return false
	}

	for i, l := 0, len(*arr1); i < l; i++ {
		if (*arr1)[i] != (*arr2)[i] {
			return false
		}
	}

	return true
}

func testSort(t *testing.T, customSortMethod sortMethod) {
	for i, l := 0, len(originalData); i < l; i++ {
		data := make([]int, len(originalData[i]))
		copy(data, originalData[i])
		customSortMethod(&data)
		if !areEqual(&data, &expectedData[i]) {
			t.Error("For", i, "Original", originalData[i], "Actual", data, "Expected", expectedData[i])
		}
	}
}

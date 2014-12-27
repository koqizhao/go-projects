
package sort

import "testing"

var data1 = [5]int { 1, 2, 3, 4, 5 }
var sortedData1 = data1

var data2 = [5]int { 5, 4, 3, 2, 1 }
var sortedData2 = data1

var data3 = [5]int { 1, 1, 1, 1, 1 }
var sortedData3 = data3

var data4 = [5]int { 1, 2, 3, 2, 1 }
var sortedData4 = [5]int { 1, 1, 2, 2, 3}

var data5 = [1]int { 1 }
var sortedData5 = [1]int { 1 }

var data6 = [0]int { }
var sortedData6 = data6

var data7 = [5]int { 1, 1, 1, 1, 2 }
var sortedData7 = data7

var data8 = [5]int { 2, 1, 1, 1, 1 }
var sortedData8 = data7

var data9 = [5]int { }
var sortedData9 = data9

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

func testSort(t *testing.T, data []int, expected []int, customSortMethod sortMethod) {
    sorted := data
    customSortMethod(&sorted)
    if !areEqual(&sorted, &expected) {
        t.Error("For", data, "Actual", sorted, "Expected", expected)
    }
}

func testSort01(t *testing.T, customSortMethod sortMethod) {
    testSort(t, data1[:], sortedData1[:], customSortMethod)
}

func testSort02(t *testing.T, customSortMethod sortMethod) {
    testSort(t, data2[:], sortedData2[:], customSortMethod)
}

func testSort03(t *testing.T, customSortMethod sortMethod) {
    testSort(t, data3[:], sortedData3[:], customSortMethod)
}

func testSort04(t *testing.T, customSortMethod sortMethod) {
    testSort(t, data4[:], sortedData4[:], customSortMethod)
}

func testSort05(t *testing.T, customSortMethod sortMethod) {
    testSort(t, data5[:], sortedData5[:], customSortMethod)
}

func testSort06(t *testing.T, customSortMethod sortMethod) {
    testSort(t, data6[:], sortedData6[:], customSortMethod)
}

func testSort07(t *testing.T, customSortMethod sortMethod) {
    testSort(t, data7[:], sortedData7[:], customSortMethod)
}

func testSort08(t *testing.T, customSortMethod sortMethod) {
    testSort(t, data8[:], sortedData8[:], customSortMethod)
}

func testSort09(t *testing.T, customSortMethod sortMethod) {
    testSort(t, data9[:], sortedData9[:], customSortMethod)
}

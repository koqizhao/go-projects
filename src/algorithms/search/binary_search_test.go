package search

import "testing"

var data = [][]int {
    []int { 1, 2, 3, 4, 5},
    []int { 1, 2, 3, 4, 5},
    []int { 1, 2, 3, 4, 5},
    []int { 1, 2, 3, 4, 5},
    []int { 1, 2, 3, 4, 5},
    []int { 1, 2, 3, 4, 5},

    []int { 1, 2, 3, 4, 5, 6},
    []int { 1, 2, 3, 4, 5, 6},
    []int { 1, 2, 3, 4, 5, 6},
    []int { 1, 2, 3, 4, 5, 6},
    []int { 1, 2, 3, 4, 5, 6},
    []int { 1, 2, 3, 4, 5, 6},
    []int { 1, 2, 3, 4, 5, 6},

    []int { 1, 1, 1, 1, 1},
    []int { 1, 1, 1, 1, 1},

    []int { 0, 0, 0, 0, 0},
    []int { 0, 0, 0, 0, 0},

    []int { 0 },
    []int { 0 },

    []int {},
}

var dataToBeSearced = []int {
    5,
    4,
    3,
    2,
    1,
    -100,

    6,
    5,
    4,
    3,
    2,
    1,
    -100,

    1,
    -100,

    0,
    -100,

    0,
    -100,

    -100,
}

var expectedResult = []int {
    4,
    3,
    2,
    1,
    0,
    -1,

    5,
    4,
    3,
    2,
    1,
    0,
    -1,

    2,
    -1,

    2,
    -1,

    0,
    -1,

    -1,
}

func TestBiSearch(t *testing.T) {
    for i, l := 0, len(data); i < l; i++ {
        isMatch := func(j int) bool {
            return data[i][j] == dataToBeSearced[i]
        }
        if result := BiSearch(len(data[i]), isMatch); result != expectedResult[i] {
            t.Errorf(`data: %v
toBeSearched: %v
actualResult: %v
expectedResult: %v
`, data[i], dataToBeSearced[i], result, expectedResult[i])
        }
    }
}


func TestBiSearchArray(t *testing.T) {
    for i, l := 0, len(data); i < l; i++ {
        if result := BiSearchArray(dataToBeSearced[i], &data[i]); result != expectedResult[i] {
            t.Errorf(`data: %v
toBeSearched: %v
actualResult: %v
expectedResult: %v
`, data[i], dataToBeSearced[i], result, expectedResult[i])
        }
    }
}

func TestBiSearchArrayRecursively(t *testing.T) {
    for i, l := 0, len(data); i < l; i++ {
        if result := biSearchArrayRecursively(dataToBeSearced[i], &data[i], 0, len(data[i])); result != expectedResult[i] {
            t.Errorf(`data: %v
toBeSearched: %v
actualResult: %v
expectedResult: %v
`, data[i], dataToBeSearced[i], result, expectedResult[i])
        }
    }
}

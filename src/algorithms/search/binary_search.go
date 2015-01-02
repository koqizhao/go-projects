package search

func BiSearchArray(data int, arr *[]int) int {
    return biSearchArray(data, arr)
}

func biSearchArray(data int, arr *[]int) int {
    for start, end := 0, len(*arr); start != end; {
        if middle := (start + end ) / 2; (*arr)[middle] == data {
            return middle
        } else if (*arr)[middle] > data {
            end = middle
        } else {
            start = middle + 1
        }
    }

    return -1
}

func biSearchArrayRecursively(data int, arr *[]int, start, end int) int {
    if start == end {
        return -1
    }

    if middle := (start + end) / 2; (*arr)[middle] == data {
        return middle
    } else if (*arr)[middle] > data {
        return biSearchArrayRecursively(data, arr, start, middle)
    } else {
        return biSearchArrayRecursively(data, arr, middle + 1, end)
    }
}

func BiSearch(arrayLength int, isMatch func(int)bool) int {
    if arrayLength < 0 {
        return -1;
    }
    return binarySearch(0, arrayLength, isMatch)
}

func binarySearch(start, end int, isMatch func(int)bool) int {
    if start == end {
        return -1
    }

    middle := (start + end) / 2
    if isMatch(middle) {
        return middle
    }

    if index := binarySearch(start, middle, isMatch); index != -1 {
        return index
    } else if index = binarySearch(middle + 1, end, isMatch); index != -1 {
        return index
    } else {
        return -1
    }
}

package sort

func QuickSort(arr *[]int) {
    if arr == nil {
        return
    }

    quickSortInternal(arr, 0, len(*arr) - 1)
}

func quickSortInternal(arr *[]int, start int, end int) {
    index := selectPivot(arr, start, end)
    if index == -1 {
        return
    }

    quickSortInternal(arr, start, index - 1)
    quickSortInternal(arr, index + 1, end)
}

func selectPivot(arr *[]int, start int, end int) int {
    if start < 0 || start >= end {
        return -1
    }

    pivot := (*arr)[start]
    for start < end {
        for ; start < end; end-- {
            if pivot > (*arr)[end] {
                (*arr)[start] = (*arr)[end]
                start++
                break
            }
        }

        for ; start < end; start++ {
            if (*arr)[start] > pivot {
                (*arr)[end] = (*arr)[start]
                end--
                break
            }
        }
    }

    (*arr)[start] = pivot
    return start
}

package sort

func QuickSort(arr *[]int) {
	if arr == nil {
		return
	}

	quickSortInternal(arr, 0, len(*arr))
}

func quickSortInternal(arr *[]int, start int, end int) {
	if start == end {
		return
	}

	index := selectPivot(arr, start, end)
	quickSortInternal(arr, start, index)
	quickSortInternal(arr, index+1, end)
}

func selectPivot(arr *[]int, start int, end int) int {
	pivot, i, j := (*arr)[start], start, end-1
	for i < j {
		for ; i < j; j-- {
			if pivot > (*arr)[j] {
				(*arr)[i] = (*arr)[j]
				i++
				break
			}
		}

		for ; i < j; i++ {
			if (*arr)[i] > pivot {
				(*arr)[j] = (*arr)[i]
				j--
				break
			}
		}
	}

	(*arr)[i] = pivot
	return i
}

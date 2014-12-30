package sort

func MergeSort(arr *[]int) {
	mergeSort(arr, 0, len(*arr))
}

func mergeSort(arr *[]int, start, end int) {
	middle := start + (end-start)/2
	if middle == start {
		return
	}

	mergeSort(arr, start, middle)
	mergeSort(arr, middle, end)

	temp := make([]int, end-start)
	index, i, j := 0, start, middle
	for ; i < middle && j < end; index++ {
		if (*arr)[i] <= (*arr)[j] {
			temp[index] = (*arr)[i]
			i++
		} else {
			temp[index] = (*arr)[j]
			j++
		}
	}
	if i == middle {
		for ; j < end; j++ {
			temp[index] = (*arr)[j]
			index++
		}
	} else if j == end {
		for ; i < start; i++ {
			temp[index] = (*arr)[i]
			index++
		}
	}

	for k := 0; k < index; k++ {
		(*arr)[start+k] = temp[k]
	}
}

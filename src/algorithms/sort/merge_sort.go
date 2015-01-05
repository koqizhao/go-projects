package sort

func MergeSort(arr *[]int) {
	if arr == nil {
		return
	}

	l := len(*arr)
	tempArr := make([]int, l)
	mergeSort(arr, 0, l, &tempArr)
	// mergeSortDownTop(arr, &tempArr)
}

func mergeSort(arr *[]int, start, end int, tempArr *[]int) {
	if end-start <= 1 {
		return
	}

	middle := (start + end) / 2
	mergeSort(arr, start, middle, tempArr)
	mergeSort(arr, middle, end, tempArr)

	for i, j, k := start, start, middle; i < end; i++ {
		if j < middle && (k >= end || (*arr)[j] <= (*arr)[k]) {
			(*tempArr)[i] = (*arr)[j]
			j++
		} else {
			(*tempArr)[i] = (*arr)[k]
			k++
		}
	}

	for i := start; i < end; i++ {
		(*arr)[i] = (*tempArr)[i]
	}
}

func mergeSortDownTop(arr *[]int, tempArr *[]int) {
	for i, l := 1, len(*arr); i < l; i *= 2 {
		for j, span := 0, i*2; j < l; j += span {
			middle, end := j+i, j+span
			if middle >= l {
				break
			}
			if end >= l {
				end = l
			}
			for k, s1, s2 := j, j, middle; k < end; k++ {
				if s1 < middle && (s2 >= end || (*arr)[s1] <= (*arr)[s2]) {
					(*tempArr)[k] = (*arr)[s1]
					s1++
				} else {
					(*tempArr)[k] = (*arr)[s2]
					s2++
				}
			}

			for k := j; k < end; k++ {
				(*arr)[k] = (*tempArr)[k]
			}
		}
	}
}

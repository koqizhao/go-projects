package sort

func MergeSort(arr *[]int) {
	l := len(*arr)
	tempArr := make([]int, l)
	mergeSort(arr, 0, l, &tempArr)
}

func mergeSort(arr *[]int, start, end int, tempArr *[]int) {
	if end - start <= 1 {
		return
	}

	middle := (start + end)/2
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

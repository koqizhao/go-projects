package sort

func MergeSort(arr *[]int) {
	mergeSort(arr, 0, len(*arr))
}

func mergeSort(arr *[]int, start, end int) {
	l := end - start
	if l < 2 {
		return
	}

	middle := (start + end)/2
	mergeSort(arr, start, middle)
	mergeSort(arr, middle, end)

	temp := make([]int, l)
	for i, j, k := 0, start, middle; i < l; i++ {
		if j < middle && (k >= end || (*arr)[j] <= (*arr)[k]) {
			temp[i] = (*arr)[j]
			j++
		} else {
			temp[i] = (*arr)[k]
			k++
		}
	}

	for i := 0; i < l; i++ {
		(*arr)[start + i] = temp[i]
	}
}

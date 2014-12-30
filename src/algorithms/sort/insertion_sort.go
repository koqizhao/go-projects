package sort

func InsertionSort(arr *[]int) {
	if arr == nil {
		return
	}

	for i, l := 1, len(*arr); i < l; i++ {
		k := 0
		for ; k < i; k++ {
			if (*arr)[k] > (*arr)[i] {
				break
			}
		}
		temp := (*arr)[k]
		(*arr)[k] = (*arr)[i]
		k++
		for ; k <= i; k++ {
			temp, (*arr)[k] = (*arr)[k], temp
		}
	}
}

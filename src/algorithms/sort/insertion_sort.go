package sort

func InsertionSort(arr *[]int) {
	if arr == nil {
		return
	}

	l := len(*arr)
	temp := make([]int, l)
	for i := 0; i < l; i++ {
		k := 0
		for ; k < i; k++ {
			if temp[k] > (*arr)[i] {
				break
			}
		}
		temp2 := temp[k]
		temp[k] = (*arr)[i]
		k++
		for ; k <= i; k++ {
			temp2, temp[k] = temp[k], temp2
		}
	}
}

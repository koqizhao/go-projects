package sort

func SelectionSort(arr *[]int) {
    if arr == nil {
        return
    }

    for i, l := 0, len(*arr) - 1; i < l; i++ {
        for j := i + 1; j <= l; j++ {
            if (*arr)[i] > (*arr)[j] {
                (*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
            }
        }
    }
}

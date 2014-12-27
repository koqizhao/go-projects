package sort

func BubbleSort(arr *[]int) {
    if arr == nil {
        return
    }

    for moved, l := false, len(*arr); ; moved = false {
        for i := 1; i < l; i++ {
            if (*arr)[i - 1] > (*arr)[i] {
                (*arr)[i], (*arr)[i - 1] = (*arr)[i - 1], (*arr)[i]
                moved = true
            }
        }

        if !moved {
            break
        }
    }
}

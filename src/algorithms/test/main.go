package main

import . "fmt"
import "os"
import "strconv"
import "algorithms/sort"

func main() {
    Println("Hello, World!");
    args := os.Args[1:]
    Println(args)

    intValues := make([]int, len(args))
    for i, l := 0, len(intValues); i < l; i++ {
        intValues[i], _ = strconv.Atoi(args[i])
    }

    intValues2 := intValues
    sort.QuickSort(&intValues2)
    Println("QuickSort: \n", intValues2)

    intValues2 = intValues
    sort.BubbleSort(&intValues2)
    Println("BubbleSort: \n", intValues2)

    intValues2 = intValues
    sort.SelectionSort(&intValues2)
    Println("SelectionSort: \n", intValues2)
}

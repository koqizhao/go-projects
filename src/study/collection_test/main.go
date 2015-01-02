package main

import (
    "fmt"
    )

func main() {
    sliceTest04()
}

// append always change the original slice, but the reference to the slice must be updated
// as the slice length changes
func sliceTest01() {
    var list []int
    pList := &list
    var rList = list
    for i := 0; i < 100; i++ {
        list = append(list, i)
        fmt.Printf(
`Original Reference:
    length: %d, capacity: %d
Pointer:
    length: %d, capacity: %d
Reference:
    length: %d, capacity: %d

`, len(rList), cap(rList), len(*pList), cap(*pList), len(list), cap(list))
    }
}

// max slice length cannot be greater than the slice capicity!
func sliceTest02() {
    var list = make([]int, 1, 2)
    fmt.Println("original length", len(list))
    list = list[:cap(list)]
    fmt.Println("after sliced", len(list))

    defer func() {
        if err := recover(); err != nil {
            fmt.Printf("max slice length cannot be greater than the slice capicity!\n\t%T \n\t%s\n", err, err)
        }
    }()
    list = list[:cap(list) + 1]
    fmt.Println("after sliced", len(list))
}

// slice is only a portion of the underlying structure
func sliceTest03() {
    list := []int { 1, 0, 3, 4, 5 }
    l1 := list[0:2]
    l1[0], l1[1] = 0, 1
    fmt.Printf("Original: %d, %d\nNew: %d, %d\n", list[0], list[1], l1[0], l1[1])
}

// slice capacity can be increased automatically
func sliceTest04() {
    list := make([]int, 1, 2)
    for i := 0; i < 65; i++ {
        list = append(list, i)
        fmt.Printf("length: %d, capacity %d\n", len(list), cap(list))
    }
}

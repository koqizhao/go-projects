package main

import "fmt"

func main() {
    fmt.Println("Hello, World!");
    var v interface{}
    var l []int
    var m map[string]int
    fmt.Printf("I: %#v, L: %v, M: %v, true: %t, %t \n", v, l, m, l == nil, m == nil);
    fmt.Printf("%d, %d\n", len(l), len(m))
    l = append(l, 1)
    // m["dd"] = 1
    fmt.Printf("%d, %d\n", len(l), len(m))

    v = l;
    v2, _ := v.([]int)
    v2[0] = 10
    fmt.Printf("%v, %v", v2[0], l[0])
}

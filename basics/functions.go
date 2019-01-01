package main

import "fmt"

// Variadic function
func sums(nums ...int) {
    fmt.Println(nums)
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func sum(a int, b int) int {
    return a + b
}

func sumsum(a, b, c int) int {
    return a + b + c
}

func vals() (int, int) {
    return 3, 5
}

func main() {
    res := sum(1, 9)
    ser := sumsum(2, 1, 9)
    fmt.Println(res)
    fmt.Println(ser)

    a, b := vals()
    fmt.Println(a, b)

    c := []int{1,2,3,4}
    sums(c...)
}

package main

import "fmt"

func main() {
    for i := 0; i < 100; i++ {
        fmt.Println(i)
    }

    // There are no whiles
    j := 0
    for j < 100 {
        fmt.Println(j)
        j++
    }
}

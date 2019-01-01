package main

import "fmt"

func main() {
    // Arrays
    var v [5]int // Arrays init with 0s
    fmt.Println("emp", v) // I can print arrays =)

    size := len(v)
    fmt.Println("size", size)

    v2 := [4]int{1, 2, 3, 4}
    fmt.Println(v2)

    var table [3][3]int
    for i := 0; i < len(table); i++ {
        for j := 0; j < len(table[0]); j++ {
            fmt.Println(table[i][j], " ")
        }
        fmt.Println()
    }

    for index1, row := range table {
        fmt.Println(index1, row)
        for index2, cell := range row {
            table[index1][index2] = 2 + cell
        }
    }

    for _, value := range table {
        fmt.Println(value)
    }

    fmt.Println(table)


    // Slices

    s := make([]string, 3)
    fmt.Println(s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    //s[3] = "d" : That would be an out of range comilation Error
    fmt.Println("get:", s[2])

    s = append(s, "d")
    fmt.Println(s)

    l := s[2:]
    fmt.Println(l)

    m := make([]string, len(s))
    copy(m, s)
    fmt.Println(m)

    b := []string{"a", "b"}
    fmt.Println(b)
}

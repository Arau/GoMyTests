package main

import "fmt"

func main() {
    m := make(map[string]int)

    m["k1"] = 3
    m["k2"] = 9

    fmt.Println(m)
    fmt.Println(len(m))

    delete(m,"k2")
    fmt.Println(m)
    fmt.Println(len(m))


    _, exists := m["k2"]
    fmt.Println(exists)

    // one liner
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println(n)

}

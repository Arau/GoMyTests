package main

import "fmt"

func main() {
    if num := 9; num < 0 { // variable num available in all the if statements
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "multiple digits")
    }
}

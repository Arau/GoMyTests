package main

import "fmt"

func main() {
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Println("I don't know the %T type\n", t)
        }
    }

    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}

package main

import "fmt"
import "time"

func worker(done chan bool) {
    fmt.Println("Working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}


func main() {

    messages := make(chan string)

    go func() { messages <- "ping" }()

    msg := <-messages // Block until both sender and reciver are ready, so I don't have to implement any sync
    fmt.Println(msg)
    fmt.Println()


    // Buffered channel
    messages = make(chan string, 2)
    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
    

    // Blocking channels
    done := make(chan bool, 1)
    go worker(done)

    <-done // Blocking until the goroutine notifies that is ready
    fmt.Println("After done")
}



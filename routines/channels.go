package main

import "fmt"
import "time"

func worker(done chan bool) {
    fmt.Println("Working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
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

    // Directional channels
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}



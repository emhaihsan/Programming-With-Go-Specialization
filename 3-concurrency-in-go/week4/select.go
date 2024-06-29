package main

import "fmt"

func main() {
    c1 := make(chan int)
    c2 := make(chan int)

    go func() {
        c1 <- 1
    }()

    go func() {
        c2 <- 2
    }()

    select {
    case msg1 := <-c1:
        fmt.Println("Received from c1:", msg1)
    case msg2 := <-c2:
        fmt.Println("Received from c2:", msg2)
    }
}

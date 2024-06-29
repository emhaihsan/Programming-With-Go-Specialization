package main

import (
    "fmt"
)

func main() {
    ch := make(chan int)

    go func() {
        fmt.Println("Sending value")
        ch <- 42
        fmt.Println("Value sent")
    }()

    value := <-ch
    fmt.Println("Received:", value)
}
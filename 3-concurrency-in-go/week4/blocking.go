package main

import "fmt"

func main() {
    c := make(chan int)

    go func() {
        c <- 42  // Mengirim data ke channel (blocking hingga data diterima)
    }()

    value := <-c  // Menerima data dari channel (blocking hingga data diterima)
    fmt.Println(value)
}

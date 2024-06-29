package main

import (
    "sync"
)

func dostuff(c1, c2 chan int, wg *sync.WaitGroup) {
    <-c1
    c2 <- 1
    wg.Done()
}

func main() {
    c1 := make(chan int)
    c2 := make(chan int)
    var wg sync.WaitGroup
    wg.Add(2)
    go dostuff(c1, c2, &wg)
    go dostuff(c2, c1, &wg)
    wg.Wait()
}

package main

import (
    "fmt"
    "sync"
)

var once sync.Once

func setup() {
    fmt.Println("Init")
}

func dostuff(wg *sync.WaitGroup) {
    once.Do(setup)
    fmt.Println("Hello")
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go dostuff(&wg)
    go dostuff(&wg)
    wg.Wait()
}

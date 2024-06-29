package main

import (
    "fmt"
    "sync"
)

var (
    counter int
    mut     sync.Mutex
)

func increment(wg *sync.WaitGroup) {
    mut.Lock()    // Mengunci sebelum mengakses shared variable
    counter++
    mut.Unlock()  // Membuka kunci setelah selesai mengakses shared variable
    wg.Done()
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go increment(&wg)
    }

    wg.Wait()
    fmt.Println("Counter:", counter)
}

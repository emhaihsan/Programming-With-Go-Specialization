package main

import (
	"fmt"
	"sync"
)

var counter int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter++
		fmt.Println("Incrementing", counter)
	}
}

func decrement(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter--
		fmt.Println("Decrementing", counter)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go increment(&wg)
	go decrement(&wg)
	wg.Wait()
	fmt.Println("Final value", counter)
}

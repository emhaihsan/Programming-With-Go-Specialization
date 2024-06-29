package main

import (
	"fmt"
	"sync"
)

func increment(wg *sync.WaitGroup, m *sync.Mutex, x *int) {
	m.Lock()
	*x++
	m.Unlock()
	wg.Done()
}

func main(){
	var x int
	var wg sync.WaitGroup
	var m sync.Mutex
	
	for i := 0; i < 100; i++ {
		wg.Add(1)
		fmt.Println("Incrementing",x)
		go increment(&wg, &m, &x)
	}
	wg.Wait()
	fmt.Println("Final value of",x)
}
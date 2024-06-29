package main 

import (
	"fmt"
	"sync"
)

func increment(wg *sync.WaitGroup,x *int) {
	*x++
	wg.Done()
}

func main(){
	var x int
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		fmt.Println("Incrementing",x)
		go increment(&wg, &x)
	}
	wg.Wait()
	fmt.Println("Final value",x)
}
package main

import (
	"fmt"
	"time"
)

func main(){
	go func(){
		fmt.Println("Goroutine started")
		time.Sleep(time.Second)
		fmt.Println("Goroutine finished")
	}()
	time.Sleep(time.Second)
	fmt.Println("Main finished")
}
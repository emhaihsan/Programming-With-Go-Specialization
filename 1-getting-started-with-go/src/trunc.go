package main

import "fmt"
func main() {
	var input float32

	fmt.Println("Enter a floating point number: ")

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
	}
	truncated := int(input)
	fmt.Println("Truncated version:", truncated)
}
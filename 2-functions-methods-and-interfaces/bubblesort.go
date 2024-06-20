package main

import "fmt"

func Swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

func BubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n; i++{
		for j := 0; j < n-i-1; j++{
			if slice[j] > slice[j+1]{
				Swap(slice, j)
		}
	}
	}
}

func main(){
	var nums []int

	fmt.Println("Enter 10 numbers: ")

	for len(nums) < 10{
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
		}
		nums = append(nums, input)
	}

	BubbleSort(nums)

	fmt.Println("Sorted: ", nums)
}
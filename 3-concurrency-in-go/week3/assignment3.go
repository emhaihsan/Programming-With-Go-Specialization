package main

import (
	"fmt"
	"sync"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func merge(parts [][]int) []int{
	switch len(parts){
	case 0:
		return nil
	case 1:
		return parts[0]
	case 2:
		return mergeTwo(parts[0], parts[1])
	default:
		mid := len(parts) / 2
		return mergeTwo(merge(parts[:mid]), merge(parts[mid:]))
	}
}

func mergeTwo(a, b []int) []int{
	result := make([]int, 0,len(a) + len(b))
	i, j := 0,0
	for i < len(a) && j < len(b){
		if a[i] < b[j]{
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

func main(){
	var n int
	fmt.Print("Enter number of integers: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Printf("Enter %d integers:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	partSize := (n +  3) / 4 // calculate size of each part, rounding up
	partitions := [][]int{
		arr[:min(partSize,n)],
		arr[min(partSize, n):min(2*partSize, n)],
		arr[min(2*partSize, n):min(3*partSize, n)],
		arr[min(3*partSize, n):n],
	}

	var wg sync.WaitGroup
	sortedParts := make([][]int, 4)

	for i := 0; i < 4; i++{
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			part := partitions[i]
			fmt.Printf("Sorted part %d: %v\n", i, part)
			sort.Ints(part)
			sortedParts[i] = part
			fmt.Printf("Sorted part %d: %v\n", i, sortedParts[i])
		}(i)
	}

	wg.Wait()

	sortedArr := merge(sortedParts)
	fmt.Printf("Sorted array: %v\n", sortedArr)
}


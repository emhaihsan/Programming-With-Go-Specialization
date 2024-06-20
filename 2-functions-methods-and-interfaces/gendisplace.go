package main

import (
	"fmt"
)

func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}

func main(){
	var a, v0, s0, t float64

	fmt.Println("Enter values for acceleration")
	fmt.Scan(&a)
	fmt.Println("Enter values for initial velocity")
	fmt.Scan(&v0)
	fmt.Println("Enter values for initial displacement")
	fmt.Scan(&s0)

	
	fn := GenDisplaceFn(a, v0, s0)
	for {
		fmt.Println("Enter time (Press '0' to quit):")
		fmt.Scan(&t)
		if t == 0 {
			break
		}
		fmt.Println("Displacement after", t, "seconds: ", fn(t))
	}


}
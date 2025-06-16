package main

import (
	"fmt" 

	"example.com/mypackage/mymath"
)

func main() {
	xs := []float64{1,2,3,4}

	avg := mymath.Average(xs)

	fmt.Println(avg)
}
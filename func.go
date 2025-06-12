package main

import (
	"fmt"

)
// ===> 1

// func average(xs []float64) float64 {
// 	var sum float64 = 0
// 	for _, v := range xs {
// 		sum += v
// 	}

// 	return sum/float64(len(xs))
// }

// func main() {
// 	xs := []float64{98,93,77,82,83}
// 	fmt.Println(average(xs))
// }


// ===> 2
/*
	By using ... before the type name of the last parame-
	ter you can indicate that it takes zero or more of those
	parameters.
*/
// func add(args ...int) int { //
// 	total := 0
// 	for _, v := range args {
// 		total += v
// 	}
// 	return total
// }
// func main() {
// 	fmt.Println(add(1,2,3))
// }

// ===> 3
// func test() {
// 	fmt.Println("hello")
// }

// func main() {
// 	x := test
// 	x()
// }

// ===> 4
// func main() {
// 	test := func (x int)  {
// 		fmt.Println(x)
// 	}

// 	test(5)
// }


// ===> 5
// func main() {
// 	test := func (x int) int  {
// 		return x*-1
// 	}(8)

// 	fmt.Println(test)
// }

// ===> 6

// func returnFanc(x string) func() {
// 	return func() {
// 		fmt.Println(x)
// 	}
// }

// func main() {
// 	returnFanc("hi")()
// 	p := returnFanc("urmi")
// 	p()

// }
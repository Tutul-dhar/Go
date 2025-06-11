package main

import (
	"fmt" 
	
)

func main() {

	// var x [5]int
	// x[4] = 100
	// fmt.Println(x)

	/*

		var total float64 = 0
		for i := 0; i < len(x); i++ {
		t	otal += x[i]
		}
		fmt.Println(total / len(x))

		The issue here is that len(x) and total have different
		types. total is a float64 while len(x) is an int. So we
		need to convert len(x) into a float64:

		fmt.Println(total / float64(len(x)))

	*/


	/*
		var total float64 = 0
		for i, value := range x {
			total += value
		}
		fmt.Println(total / float64(len(x)))

		In this for loop i represents the current position in the
		array and value is the same as x[i]. We use the keyword range followed by the name of the variable we
		want to loop over.

		The Go compiler won't allow you to create variables
		that you never use. Since we don't use i inside of our
		loop we need to change it to this:

		var total float64 = 0
		for _, value := range x {
			total += value
		}
		fmt.Println(total / float64(len(x)))

	*/


	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i
	}

	ar1 := arr[1:4]
	fmt.Println(ar1)

}
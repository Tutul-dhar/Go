package main

import "fmt"

func main() {
 	/*var name string = "Tutul dhar"
	fmt.Println(name)
	*/

	// var x string
 	// x = "Hello World"
 	// fmt.Println(x)

	// var x string
	// x = "first"
	// fmt.Println(x)
	// x = "second"
	// fmt.Println(x)

	// var x string
	// x = "first "
	// fmt.Println(x)
	// x = x + "second"
	// fmt.Println(x)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)

	/*

		var name string = "Tutul"
		Breakdown:
		var → keyword to declare a variable

		name → the variable name

		string → type (e.g., string, int, float64)

		"Tutul" → value
	*/
}
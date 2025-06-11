package main

import "fmt"

func main() {
	/*
		var x map[string]int

		The map type is represented by the keyword map, followed by the key type in brackets and finally the value
		type. If you were to read this out loud you would say “x
		is a map of strings to ints.

	*/
	/*
		Creating a Map
		Option 1: make() function
		ages := make(map[string]int)

		Option 2: Literal syntax
		ages := map[string]int{
			"Alice": 30,
			"Bob":   25,
		}

	*/
		
	/*
		Accessing Map Values
		fmt.Println(ages["Alice"]) // Output: 30
	*/


	/*
		Deleting from a Map
		delete(ages, "Alice")

	*/

	/*
		Check if Key Exists
		value, exists := ages["Alice"]

		if exists {
			fmt.Println("Age:", value)
		} else {
			fmt.Println("Not found")
		}

	*/

	/*
		Looping through a Map
		for key, value := range ages {
    		fmt.Println(key, "is", value, "years old")
		}

	*/

	//var ma map[string]int
	ma := make(map[string]int)
	ma["tutul"] = 10

	fmt.Println(ma["tutul"])
}
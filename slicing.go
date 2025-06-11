package main

import "fmt"

func main() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i
	}

	ar1 := arr[1:4] // print all value that are in index >=low and < high
	fmt.Println(ar1)

	/*
		Omitting start or end
		nums := []int{10, 20, 30, 40, 50}

		fmt.Println(nums[:3]) // [10 20 30]
		fmt.Println(nums[2:]) // [30 40 50]
		fmt.Println(nums[:])  // [10 20 30 40 50]

	*/

	/*
		Slice Is a Reference

		original := []int{1, 2, 3, 4}
		part := original[1:3]
		part[0] = 100

		fmt.Println(original) // [1 100 3 4]

	*/

	/*
		Using append() with slices
		nums := []int{1, 2}
		nums = append(nums, 3, 4)

		fmt.Println(nums) // [1 2 3 4]

	*/

	/*
		Creating slices with make()
		s := make([]int, 3) // length 3, filled with zeroes
		fmt.Println(s)      // [0 0 0]

		s := make([]int, 3, 5) // len 3, cap 5

	*/

	/*
		func main() {
			slice1 := []int{1,2,3}
			slice2 := make([]int, 2)
			copy(slice2, slice1)
			fmt.Println(slice1, slice2)
		}

		After running this program slice1 has [1,2,3] and
		slice2 has [1,2]. The contents of slice1 are copied
		into slice2, but since slice2 has room for only two elements only the first two elements of slice1 are copied.

	*/
}
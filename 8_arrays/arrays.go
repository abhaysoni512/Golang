package main

import "fmt"

// numbered sequence of specefic length
func main()  {
	// zero valued :
	// int --> 0, string --> "" , bool --> false


	var nums[10]int // declare array length

	// array length
	//fmt.Println(len(nums))

	nums[0] = 1
	// fmt.Println(nums[0])

	// printing whole array
	fmt.Println(nums) // Not defined values are fall into 0


	// Adding elements while declaring
	nums1 := [3]int{1,2,3}
	fmt.Println(nums1)


	// declearing 2d array

	nums2d := [3][3]int{{1,2,3},{4,5,6}}
	fmt.Println(nums2d)

}
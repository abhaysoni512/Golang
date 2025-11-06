package main

import (
	"fmt"
	"slices"
)

// slice --> dynamic
// most used construct in go

func main()  {
	//uninitilised slice in nill
	var nums[]int
	fmt.Println(nums == nil)
    // Adding elements to slice
	nums = append(nums, 11)
	nums = append(nums, 22)
	nums = append(nums, 32)
	fmt.Println(nums)
	fmt.Println(cap(nums))
	
	// copy function
	nums1 := make([]int, 2)  // (slice, initilal_size, capacity)
	nums1[0] = 1
	nums1[1] = 2
	fmt.Println("before copy ", nums1)
	copy(nums1, nums)  // copy(dest, source) // here only first two elements will be copied
	fmt.Println("after copy ", nums1)

	// slice operator
	var nums3 = []int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(nums3[0:3]) // 0 to 2 index
	fmt.Println(nums3[3:6]) // 3 to 5 index
	fmt.Println(nums3[:6])  // 0 to 5 index
	fmt.Println(nums3[3:])  // 3 to end index

	// slices package
	var num4 = []int{1,2,3}
	var num5 = []int{4,5,6}
	fmt.Println(slices.Equal(num4, num5)) // false


	println("-------------------")
	// 2d slice
	nums2d := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println(nums2d)
	
	
}
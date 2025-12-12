package main

import (
	"fmt"
)

// TC: o(n)  sc: o(n)
func DeleteDuplicate(arr []int) []int {
	// Brute force : sort , 1 2,2,2,3,3,4,5,5
	// we will use map
	mp := make(map[int]struct{})

	for _, val := range arr {
		_, ok := mp[val]
		if ok {
			continue
		} else {
			mp[val] = struct{}{}
		}
	}
	var RemoveDuplicate []int
	for ele := range mp {
		RemoveDuplicate = append(RemoveDuplicate, ele)
	}
	return RemoveDuplicate
}

// TC : O(N) & SC : O(1) use in case of sorted array

func RemoveDuplicates(arr []int) {
	index := 1
	for i := 1; i <= len(arr)-1; i++ {
		if arr[i] != arr[i-1] {
			arr[index] = arr[i]
			index++
		}
	}
	arr = arr[:index]

	fmt.Println(arr)
}

// Given an array arr[] of size n, the task is to find all the Leaders in the array. An element is a Leader if it is greater than or equal to all the elements to its right side o(n^2) , o(n)
func LeaderInArray(arr []int) []int {

	var leaderArr []int

	for i := 0; i < len(arr); i++ {
		j := i + 1
		for ; j < len(arr); j++ {
			if arr[i] < arr[j] {

				break
			}
		}
		fmt.Println(j)
		if j == (len(arr)) {
			leaderArr = append(leaderArr, arr[i])
		}
	}
	return leaderArr

}

// optimized :Suffix Maximum , O(N) , O(N) The idea is to scan all the elements from right to left in an array and keep track of the maximum till now. When the maximum changes its value, add it to the result. Finally reverse the result
func LeaderInArrayAdv(arr []int) []int {

	var leaderArr []int
	n := len(arr)
	maxRight := arr[n-1]

	leaderArr = append(leaderArr, maxRight)

	for i := n - 2; i >= 0; i-- {
		if arr[i] > maxRight {
			maxRight = arr[i]
			leaderArr = append(leaderArr, maxRight)
		}
	}
	return leaderArr
}

func main() {
	//DupArr := []int{1, 2, 2, 2, 3, 3, 4, 5, 5}
	// fmt.Println(DeleteDuplicate(DupArr))
	// DeleteDuplicate(DupArr)
	// fmt.Println(DupArr)
	// RemoveDuplicates(DupArr)
	leadArr := []int{16, 17, 4, 3, 5, 2}
	//fmt.Println(LeaderInArray(leadArr))
	fmt.Println(LeaderInArrayAdv(leadArr))

}

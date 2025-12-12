// Select minimum and swap

// Algorithm :

// 1. Get minimum element in list {13, 46, 24, 52, 20, 9}
// 2. once got swap with current pointer selected elements {9, 46, 24, 52, 20 , 13}
// 3. Now, as 9 is at right position, move to sub array window of 46 to 13 , get minimum do repeat the same

package main

import "fmt"



func SelectionSort(arr []int) {
	for i := 0; i <= len(arr)-2; i++ {
		mini_index := i
		for j := i; j <= len(arr)-1; j++ {
			if arr[j] < arr[mini_index] {
				mini_index = j
			}
		}
		arr[mini_index], arr[i] = arr[i], arr[mini_index]
	}
}

// Push the max elements to last by adjust swapping
func BubbleSort(arr []int){
	var IsSorted bool
	for i := len(arr)-1; i>=1;i--{
		IsSorted = false
		for j := 0 ; j<=i-1 ;j++{
			if arr[j] > arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
				IsSorted = true
			}
		}
		if IsSorted == false {
			fmt.Println("Array is already sorted")
			return
		}
	
	}
}

// takes an element & place it in correct order
// 14 9 15 12 6 8 13
// Algorithm : Imagine : 14 as window is it at correct position, yes
// Algorithm : 14 9 : Is 9 at correct position, no , place it at correct position
// Algorithm : 9 14 15 : Is 15 at correct position, yes
// Algorithm : 9 14 15 12 : Is 12 at correction positon , No, keep swaping till left element is bigger than right
func InsertionSort(arr []int){
	for i :=1;i<= len(arr)-1; i++{
		
		for j :=i ; j>0 ;j--{
			if arr[j-1]>arr[j] {
				arr[j-1],arr[j] = arr[j], arr[j-1]
			}
			
		}
		
	}
}

func main() {
	arr := []int{11 , -1 ,0 ,2,-5}
	fmt.Println("Before swapping : ", arr)
	//SelectionSort(arr)
	InsertionSort(arr)
	fmt.Println("After swapping : ", arr)
}

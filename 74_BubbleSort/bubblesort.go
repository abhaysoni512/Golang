// Push the maximum to the last by adjacent swapping

package main

import "fmt"

func swap(a, b *int) {
	tmp := 0
	tmp = *a
	*a = *b
	*b = tmp
}

func BubbleSort(arr []int) {
	flag_swap := false
	for i := 0; i < len(arr)-1; i++ {

		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
				flag_swap = true
			}
		}
		if flag_swap == false {
			fmt.Println("No swapping required")
			return
		}
	}
}

func main() {
	arr := []int{11, 22, 33, 44, 55}
	fmt.Println("Before swapping : ", arr)
	BubbleSort(arr)
	fmt.Println("After swapping : ", arr)

}

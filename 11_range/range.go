package main

import "fmt"

func main()  {
	// nums := []int{6,7,8}
	// for i:=0 ;i<3 ;i++{
	// 	fmt.Print(nums[i]," ")
	// }

	// sum : 
	// sum :=0
	// for i, num := range nums{ // i = index
	// 	fmt.Println(i)
	// 	sum += num

	// }
	// fmt.Println(sum)

	// iteration over map
	// m := map[string]int{"price ":40, "buy ":300 }
	// for key,value := range m{ // if we take only single value it will return key
	// 	fmt.Println(key,":",value)
	// }


	// unicode code point rune
	// starting byte of rune
	// 255--> 1 byte 
	// 300--> 2 byte 
	// suppose we have first char that fits in 1 byte so 0, but let's say we got 2 byte char so then element will as 
	// 3 :-       ex. 0  103      |  2 111
	for i, c := range "golang"{
		fmt.Println(i,c)
	}
	

	
}

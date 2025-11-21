package main

import (
	"fmt"
	"strings"
)
type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool    { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int)       { s[i], s[j] = s[j], s[i] }
func main(){
	// array := [5]int{1,2,3,4,5}
	// slice := array[1:4]

	// for i := range 5{
	// 	fmt.Printf("Array %d having address %p\n",i,&array[i])
	// }
	// fmt.Println("__________Intial__________________")
	// fmt.Println("Array : ",array)
	// fmt.Println("Slice : ",slice)
	// fmt.Println("Slice initial addresss : ",&slice[0])

	// fmt.Println("Initial len , Inital capacity of slice ",len(slice),cap(slice))

	// fmt.Println("__________Modification__________________")
	// slice[0] = -1
	// fmt.Println("Array : ",array)
	// fmt.Println("Slice : ",slice)

	//var arr[5] int

  	// arr := [5]int{1,2,3,4,5}

	// my_slice1 := arr[0:] // 1 2 3 4 5
	// fmt.Println(my_slice1)


	// var make_slice = make([]int,2,4)
	// my_slice2 := []string{"This", "is", "the", "tutorial",
    //     "of", "Go", "language"}

	//1st way of iteration 
	// for i:=0;i<len(my_slice2);i++{
	// 	fmt.Print(my_slice2[i]," ")
	// 	if i == len(my_slice2)-1{
	// 		fmt.Println()
	// 	}
	// }

	//2nd range based

	// for _,val := range my_slice2{
	// 	fmt.Println(val)
	// }
	


	// sort a slice :

	 

	// mySlice := []int{-2,3,42,33}

	// sort.Sort(IntSlice(mySlice))

	// fmt.Println("mySlice : ",mySlice)
	

	 // func Trim(ori_slice[]byte, cut_string string) []byte


	//  slice1 := []byte{'!', '!', 'G', 'e', 'e', 'k', 's', 'f',  
    //              'o', 'r', 'G', 'e', 'e', 'k', 's', '#', '#'}

	// fmt.Printf("%s\n",slice1)
	
	// res1 := bytes.Trim(slice1,"!#")
	// fmt.Printf("%s\n",res1)


	// creation 

	// strings.Trim removes all leading and trailing characters from the string s that are present in the cutset string.
	s := "@@Hello, Geeks!!"

	results := strings.Trim(s,"")

	fmt.Println(results)



	 

	
	
}
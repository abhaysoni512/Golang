package main
func varadic_sum_fun(nums ...int) int{
	total := 0
	for _, num := range nums{
		total+= num
	}
	return total
}
func main()  {
	varadic_sum_fun(1,2,3,4,5,6,7,8,9,10)
	println(varadic_sum_fun(1,2,3,4,5,6,7,8,9,10))
	
	
	// how to varadic function with slice
	nums := []int{1,2,3,4,5,6,7,8,9,10}
	sum := varadic_sum_fun(nums...)  //need to use ... with slice
	println(sum)

}
package main


//it will generate infinte stream of data for channel based on some function passed to it
func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T{
	
}
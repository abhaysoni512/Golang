1. Rune in Go lang : 
    A rune is data type in go lang which is superset of ASCII and represents a Unicode code point.
    It is an alias for int32 and can represent any character in the Unicode standard.

2. Defer in Go lang :
    The defer statement in Go is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup.
    Deferred functions are executed in LIFO order just before the surrounding function returns.    
    defer function or method call arguments evaluate instantly, but they don't execute until the nearby functions returns
    Defer statements are generally used to ensure that the files are closed when their need is over, or to close the channel, or to catch the panics in the program.

3. Slice in Go lang : 
    A slice in go is data structure that provides a dynamic, flexible view into the elements of an array.
    Slices are more powerful than arrays as they can grow and shrink in size.
    A slice is a reference to a portion of an array and consists of three components: a pointer to the underlying array, the length of the slice, and its capacity.
    Slices are having three components: 
        1. Pointer: It points to the first element of the array that is accessible through the slice.. Here, it is not necessary that the pointed element is the first element of the array.
        2. Length:The length is the total number of elements present in the array.
        3. Capacity: It represents the maximum number of elements that can be stored in the slice without allocating more memory.

    Creation of slice:
        1. Using slice literal
        var my_slice_1 = []string{"Geeks", "for", "Geeks"}
        2. Using an Array
        arr := [4]string{"Geeks", "for", "Geeks", "GFG"}

            // Creating slices from the given array
            var my_slice_1 = arr[1:2]

            // len: 1, cap: 3
        3. Using make() function
        my_slice := make([]int, 5, 10)

4. Go rotuines in Go lang :
    A goroutine is a lightweight thread of execution in the Go programming language.
    Goroutines are managed by the Go runtime, which handles their scheduling and execution.
    They are created using the go keyword followed by a function call.
    When a goroutine is created, it runs concurrently with other goroutines in the same program.
    Goroutines communicate with each other using channels, which provide a way to send and receive data between them safely.
    Goroutines are designed to be lightweight and efficient, allowing developers to create thousands or even millions of them in a single program without incurring significant overhead.
5. Channels in Go lang :
    Channels are a powerful feature in Go that facilitate communication and synchronization between goroutines.
    They provide a way for goroutines to send and receive values of a specified type.
    Channels are created using the make() function, specifying the type of data they will carry.
    There are two types of channels: unbuffered and buffered.
    Unbuffered channels require both the sender and receiver to be ready at the same time for communication to occur, while buffered channels allow sending multiple values without an immediate receiver, up to the specified buffer size.
    Channels support operations such as sending (using the <- operator), receiving (also using the <- operator), and closing (using the close() function).
    Closing a channel indicates that no more values will be sent on it, allowing receivers to handle the end of communication gracefully.
    Channels can also be used in select statements, enabling goroutines to wait on multiple channel operations simultaneously.

5. waitgroup in Go lang :
    A WaitGroup in Go is a synchronization primitive provided by the sync package that allows you to wait for a collection of goroutines to finish executing.
    It is used to manage concurrent operations and ensure that the main program or other goroutines can wait for a set of goroutines to complete before proceeding.
    A WaitGroup maintains a counter that tracks the number of active goroutines.
    You can increment the counter using the Add() method before starting a new goroutine and decrement it using the Done() method when a goroutine completes its work.
    The Wait() method blocks the calling goroutine until the counter reaches zero, indicating that all goroutines have finished executing.
    This mechanism helps prevent premature termination of the main program and ensures proper synchronization between concurrent tasks.
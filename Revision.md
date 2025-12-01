1. Rune in Go lang : 
    A rune is data type in go lang which is superset of ASCII and represents a Unicode code point.
    It is an alias for int32 and can represent any character in the Unicode standard.

2. Defer in Go lang :
    The defer statement in Go is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup.They are executed after the surrounding function returns, in the reverse order they were deferred.
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
    A goroutine is a lightweight thread managed by the Go runtime that allows functions (or methods) to run concurrently with other goroutines, all sharing the same memory address space.
    "in the same address space" :-

    Unlike operating system threads or processes (which often have separate memory spaces), all goroutines in a Go program share the same memory address space.
        This means:
        They can directly access the same variables (no need for message passing like in processes).
        But you must be careful with race conditions → that's why Go gives you channels and sync primitives.

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

6. Composition in Go lang :
    Composition in Go is a design principle that allows you to build complex types by combining simpler types.
    In Go, composition is typically achieved by embedding one struct type within another struct type.
    This allows the outer struct to inherit the fields and methods of the embedded struct, promoting code reuse and modularity.


7 Go is object oriented lang ?
    Go is not a traditional object-oriented programming (OOP) language like Java or C++, but it does support some OOP concepts.
    Go does not have classes or inheritance, which are fundamental features of traditional OOP languages.
    Instead, Go uses structs to define custom data types and methods to associate behavior with those types.
    Go supports encapsulation  through the use of exported and unexported fields and methods, allowing you to control access to the internal state of a struct. exported fields and methods (those that start with an uppercase letter) are accessible from other packages, while unexported ones (those that start with a lowercase letter) are not.
    Encapsulation helps in bundling the data and methods that operate on that data within a single unit.
    Polymorphism is achieved in Go through interfaces, which define a set of method signatures that a type must implement to satisfy the interface.
    Inheritance is not supported in Go, but composition allows you to build complex types by combining simpler types.
    Abstraction is supported through interfaces, which allow you to define behavior without specifying the underlying implementation.
    Overall, while Go does not fully embrace traditional OOP principles, it provides enough features to enable developers to use OOP concepts in a way that fits the language's design philosophy.

7. Varadic functions in Go lang :
    Variadic functions in Go are functions that can accept a variable number of arguments.
    They are defined by using an ellipsis (...) before the type of the last parameter in the function signature.
    This allows you to pass zero or more arguments of that type when calling the function.
    Inside the function, the variadic parameter is treated as a slice of the specified type, allowing you to iterate over the arguments or access them using indexing.
    Variadic functions are useful when you want to create functions that can handle a flexible number of inputs without requiring the caller to explicitly create a slice or array.

8. Defer vs Panic vs Recover in Go lang :
    Defer is used halt the execution of underlying function until the surrounding function returns.
    Panic is used to stop the normal execution of the program when an unexpected error occurs.
    Recover is used to regain control of a panicking goroutine. It is used inside a deferred function to catch a panic and prevent the program from crashing.

9. Interfaces & Polymorphism in Go lang :
    An interface in Go defines a set of method signatures that a type must implement to satisfy the interface.
    Polymorphism is the  ability of a VALUE of certain TYPE to also be another TYPE.
    In Go, values can be more than one type if they implement multiple interfaces.

10. Closures in Go lang :
    A closure in Go is a function that captures variables from its surrounding scope and keeps them alive even the scope has finished executing.
    Closures are created when a function is defined inside another function and references variables from the outer function.

11. Package atomic

    The atomic package in Go provides low-level atomic memory primitives for synchronizing access to shared variables across multiple goroutines.
    It offers functions for performing atomic operations on integers and pointers, ensuring that these operations are executed atomically, without interruption.
    This is crucial for preventing race conditions when multiple goroutines read from and write to the same variable concurrently.

    Difference between sync.Mutex and atomic package:
    1. Granularity:
        - sync.Mutex: Provides coarse-grained locking, allowing you to lock and unlock a section of code, which can lead to contention if many goroutines are trying to access the same resource.
        - atomic package: Offers fine-grained atomic operations on individual variables, reducing contention and improving performance for simple operations.
    2. Use Cases:
        - sync.Mutex: Suitable for protecting complex data structures or critical sections of code where multiple operations need to be performed atomically.
        - atomic package: Ideal for simple operations on single variables, such as incrementing a counter or swapping pointers, where the overhead of a mutex is unnecessary.

11. Fan in and Fan out in Go lang :
    
    Fan-In is a concurrency pattern where multiple goroutines send their results to a single output channel. This pattern is useful for collecting results from various concurrent operations into one place for further processing or aggregation.


    Fan-Out is a concurrency pattern where multiple worker goroutines read from a single input channel to process data concurrently. This pattern helps to distribute the workload and speed up processing by utilizing multiple goroutines.

12. Context in Go lang :
    Context in Go is a tool that we can use with concurrent design pattern to make sure that if we have like a process which get lost or cancelled we can propagate that cancellation to all the goroutines which are working on that process.

    Context could also help us with like passing arround variables which are request scoped like auth tokens, user ids etc.

    Real time example : In Go servers, each incoming request is handled in its own goroutine. Request handlers often start additional goroutines to access backends such as databases and RPC services. The set of goroutines working on a request typically needs access to request-specific values such as the identity of the end user, authorization tokens, and the request’s deadline. When a request is canceled or times out, all the goroutines working on that request should exit quickly so the system can reclaim any resources they are using. At Google, we developed a context package that makes it easy to pass request-scoped values, cancellation signals, and deadlines across API boundaries to all the goroutines involved in handling a request. The package is publicly available as context. This article describes how to use the package and provides a complete working example.

    Request-scoped values (also called request-scoped context or per-request context) are pieces of data that live only for the duration of a single incoming request and are automatically cleaned up when the request ends.

    Why "across API boundaries" matters
    In modern systems (especially microservices or clean/hexagonal architecture), a single incoming HTTP/gRPC request often crosses many boundaries:

    HTTP Request 
    → API Gateway / Edge service 
        → Authentication service (internal API call)
        → User service (another internal call)
        → Billing service 
        → Database layers, logging, auditing, etc.

    You usually need certain data available everywhere during that request, for example:

        User ID / tenant ID (after authentication)
        Trace/correlation ID (for distributed tracing)
        Request locale / timezone
        Authorization decisions or roles
        Transaction/context flags (dry-run mode, impersonation, etc.)
    These mechanisms give you a map/dictionary that:

        Is automatically created when a request starts
        Is automatically destroyed when the request finishes
        Is accessible from anywhere in the call stack (even deep in libraries)
        Does not leak to other concurrent requests (very important for thread-safety)
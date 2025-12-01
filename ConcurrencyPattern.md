1. for - select loop

Scenerio: chars = ['a','b','c','d','e'], we want to send these chars to a channel one by one .
```
    for char := range chars{
        select{
            case channel<- char:
                fmt.Println("sent ", char)
        }
    }
```
2. done channel:
    done channel pattern : Here parent goroutine signals to child goroutine to stop its work, it uses a done channel to send the signal.

3. pipelines

    What is pipeline?
    A pipeline is a series of processing stages where the output of one stage is the input to the next stage. In Go, pipelines are often implemented using goroutines and channels to facilitate concurrent processing.


4. worker pool

    A worker pool is a concurrency pattern where a fixed number of worker goroutines are created to process tasks from a shared job queue. This pattern helps to limit the number of concurrent tasks being processed, which can be useful for managing resource usage and improving performance.
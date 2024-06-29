# Programming with Go Specialization

This repository contains codes and tasks from Specialization Programming with Go. Each module has its own directory that covers a variety of topics from an introduction to Go to advanced techniques such as goroutines and synchronization.

## Directory Structure

- **1-getting-started-with-go**
 - **names.txt**: Simple text file.
 - **read.go**: Example of reading user input.
 - **slice.go**: Example of using slices in Go.
 - **trunc.go**: Example of using the `trunc` function in Go.

- **2-functions-methods-and-interfaces**
 - **animal-interface.go**: Examples of using interfaces in Go.
 - **animal.go**: Simple implementation of the `Animal` interface.
 - **bubblesort.go**: Implementation of the Bubble Sort algorithm.
 - **gendisplace.go**: Example of generic displacement implementation.

- **3-concurrency-in-go**
 - **week2**
 - **assignmentweek2.go**: The second week's assignment is about concurrency.
 - **interleavings.go**: Examples of interleavings in concurrency.
 - **mutex.go**: Example of using mutex for synchronization.
 - **raceconditions.go**: Examples of race conditions in concurrency programs.
 - **threads.go**: Example of using threads in Go.

 - **week3**
 - **assignment3.go**: Third week assignment about concurrency.
 - **basicsync.go**: Example of basic sync in Go.
 - **blockinggoroutines.go**: Example of a blocked goroutine.
 - **buffer.go**: Example of using buffers in Go.
 - **communication.go**: Example of communication between goroutines.
 - **exitgoroutine.go**: Example of how to exit a goroutine.
 - **goroutines.go**: Basic examples of using goroutines.
 - **waitgroup.go**: Example of using WaitGroup for goroutine synchronization.

 - **week4**
 - **assignmentweek4.go**: Fourth week assignment about concurrency.
 - **blocking.go**: Example of blocking in a goroutine.
 - **deadlock.go**: Examples of deadlock situations and how to avoid them.
 - **dining.go**: Dining Philosophers issue implementation.
 - **mutex.go**: Advanced example of using mutex.
 - **oncesync.go**: Example of using sync.Once for a one-time sync.
 - **select.go**: Example of using select statements in Go.

## Use

Each directory and file has code examples and tasks that can be run independently. To run each code example, use the following command:

```bash
go run <file_name>.go
```

For example, to run an example using slice:

```bash
go run 1-getting-started-with-go/slice.go
```

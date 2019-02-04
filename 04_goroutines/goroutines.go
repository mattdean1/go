// Homework 5: Goroutines
// Due March 3, 2017 at 11:59pm
package main

import (
	"fmt"
	// "errors"
)

func main() {
	// Feel free to use the main function for testing your functions
	hello := make(chan string, 5)
	hello <- "Hello world"
	hello <- "ÐŸÑ€Ð¸Ð²ÐµÑ‚ Ð¼Ð¸Ñ€"
	hello <- "ÐŸÑ€Ð¸Ð²Ñ–Ñ‚ Ð¡Ð²Ñ–Ñ‚"
	hello <- "Witaj Å›wiecie"
	close(hello)
	fmt.Println(<-hello)
	for greeting := range hello {
		fmt.Println(greeting)
	}
}

// Filter copies values from the input channel into an output channel that match the filter function p
// The function p determines whether an int from the input channel c is sent on the output channel
func Filter(input <-chan int, predicate func(int) bool) (<-chan int, error) {
	output := make(chan int, cap(input))
	for inputValue := range input {
		if predicate(inputValue) {
			output<-inputValue
		}
	}
	close(output)
	return output, nil
	// errors.New("fail")
}

// Result is a type representing a single result with its index from a slice
type Result struct {
	index  int
	result string
}

// ConcurrentRetry runs all the tasks concurrently and sends the output in a Result channel
//
// concurrent is the limit on the number of tasks running in parallel. Your
// solution must not run more than `concurrent` number of tasks in parallel.
//
// retry is the number of times that the task should be attempted. If a task
// returns an error, the function should be retried immediately up to `retry`
// times. Only send the results of a task into the output channel if it does not error.
//
// Multiple instances of ConcurrentRetry should be able to run simultaneously
// without interfering with one another, so global variables should not be used.
// The function must return the channel without waiting for the tasks to
// execute, and all results should be sent on the output channel. Once all tasks
// have been completed, close the channel.
func ConcurrentRetry(tasks []func() (string, error), concurrent int, retry int) <-chan Result {
	// TODO
	return nil
}

// Task is an interface for types that process integers
type Task interface {
	Execute(int) (int, error)
}

// Fastest returns the result of the fastest running task
// Fastest accepts any number of Task structs. If no tasks are submitted to
// Fastest(), it should return an error.
// You should return the result of a Task even if it errors.
// Do not leave any pending goroutines. Make sure all goroutines are cleaned up
// properly and any synchronizing mechanisms closed.
func Fastest(input int, tasks ...Task) (int, error) {
	// TODO
	return 0, nil
}

// MapReduce takes any number of tasks, and feeds their results through reduce
// If no tasks are supplied, return an error.
// If any of the tasks error during their execution, return an error immediately.
// Once all tasks have completed successfully, return the value of reduce on
// their results in any order.
// Do not leave any pending goroutines. Make sure all goroutines are cleaned up
// properly and any synchronizing mechanisms closed.
func MapReduce(input int, reduce func(results []int) int, tasks ...Task) (int, error) {
	// TODO
	return 0, nil
}
package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int) {
	for {
		fmt.Printf("Worker %d - reading...\n", id)
		v := <-tasks
		fmt.Printf("Worker %d - processed %d!\n", id, v)
		time.Sleep(time.Millisecond)
		fmt.Printf("Worker %d - sending result %d...\n", id, v)
		results <- v
		fmt.Printf("Worker %d - sent result %d!\n", id, v)
	}
}

func producer(n int, tasks chan<- int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Task %d sending...\n", i)
		tasks <- i
		fmt.Printf("Task %d sent!\n", i)
	}
}

func main() {
	var n = 100
	var ncpu = runtime.NumCPU()

	// Tried different buferization 0, 1, ncpu, ncpu * 2
	// - 0 implies goroutine context switch on each channel send
	// - 1 only one worker gets task, other have to block, similar to 0
	// - ncpu looks like the most promising with minimum of switches
	// - further increase doesn't help, since producer is faster than workers
	//   and just fills channel to the top while workers read ncpu tasks at max
	tasks := make(chan int, ncpu)
	results := make(chan int, ncpu)

	for i := 0; i < ncpu; i++ {
		go worker(i, tasks, results)
	}

	// Necessary to run in separate goroutine to avoid deadlocking
	go producer(n, tasks)

	for i := 0; i < n; i++ {
		res := <-results
		fmt.Println(res)
	}
}

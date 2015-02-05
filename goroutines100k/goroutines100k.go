package main

import (
	"flag"
	"fmt"
	"github.com/alexgula/go-playground/timeit"
	"runtime"
	"runtime/debug"
)

var ngoroutine = flag.Int("n", 100000, "how many")

func f(left, right chan int) {
	left <- 1 + <-right
}

func runSeq() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0
	x := <-leftmost
	fmt.Println(x)
}

func runParallel() {
	flag.Parse()
	left, right := make(chan int), make(chan int)
	for i := 0; i < *ngoroutine; i++ {
		go f(left, right)
	}
	x := 0
	for i := 0; i < *ngoroutine; i++ {
		right <- 0 // bang!
		x += <-left
	}
	fmt.Println(x)
}

func runParallelFast() {
	flag.Parse()
	left, right := make(chan int), make(chan int)
	x := 0
	for i := 0; i < *ngoroutine; i++ {
		go f(left, right)
		right <- 0 // bang!
		x += <-left
	}
	fmt.Println(x)
}

func main() {
	fmt.Printf("GOMAXPROCS was %v\n", runtime.GOMAXPROCS(runtime.NumCPU()))
	//fmt.Printf("GOMAXPROCS now %v\n", runtime.GOMAXPROCS(-1))
	debug.SetGCPercent(-1)
	timeit.RunFmt(runSeq, 1)
	timeit.RunFmt(runParallel, 1)
	timeit.RunFmt(runParallelFast, 1)
}

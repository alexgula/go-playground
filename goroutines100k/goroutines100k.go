package goroutines100k

func f(left, right chan int) {
	left <- 1 + <-right
}

func runSeq(ngoroutine int) int {
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0
	x := <-leftmost
	return x
}

func runParallel(ngoroutine int) int {
	left, right := make(chan int), make(chan int)
	for i := 0; i < ngoroutine; i++ {
		go f(left, right)
	}
	x := 0
	for i := 0; i < ngoroutine; i++ {
		right <- 0 // bang!
		x += <-left
	}
	return x
}

func runParallelFast(ngoroutine int) int {
	left, right := make(chan int), make(chan int)
	x := 0
	for i := 0; i < ngoroutine; i++ {
		go f(left, right)
		right <- 0 // bang!
		x += <-left
	}
	return x
}

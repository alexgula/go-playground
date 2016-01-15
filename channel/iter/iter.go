package iter

func Range(n int) (result chan int) {
	result = make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			result <- i
		}
		close(result)
	}()
	return
}

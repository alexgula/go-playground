package merge

func merge(out chan<- int, a, b <-chan int) {
	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil
				continue
			}
			out <- v
		case v, ok := <-b:
			if !ok {
				b = nil
				continue
			}
			out <- v
		}
	}
	close(out)
}

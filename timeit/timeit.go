package timeit

import (
	"fmt"
	"time"
)

func Run(f func(), times int) time.Duration {
	timeStart := time.Now()
	for i := 0; i < times; i++ {
		f()
	}
	return time.Since(timeStart)
}

func RunFmt(f func(), times int) {
	fmt.Printf("Duration: %v msec", Run(f, times).Seconds()*1000)
	if times > 1 {
		fmt.Printf(", count: %v\n", times)
	}
	fmt.Println()
}

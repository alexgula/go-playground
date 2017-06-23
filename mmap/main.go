package main

import mmap "github.com/edsrzf/mmap-go"
import "os"
import "fmt"
import "time"

func main() {
	var valueCount = 32 * 1024 * 1024 * 1024
	var valueSize = 1
	var capacity = valueCount * valueSize
	var trackFrequency = 256 * 1024 * 1024

	var start = time.Now()

	f, err := os.OpenFile("large.data", os.O_RDWR|os.O_CREATE, 0600)
	assert(err)
	defer f.Close()

	f.WriteAt([]byte{0}, int64(capacity-1))

	var elapsed = time.Since(start)

	fmt.Printf("File created in %v\n", elapsed)

	m, err := mmap.MapRegion(f, capacity, mmap.RDWR, 0, 0)
	assert(err)
	defer m.Unmap()

	start = time.Now()

	for i := 0; i < capacity; i++ {
		m[i] = 0
		if i%trackFrequency == 0 {
			fmt.Print(".")
		}
	}
	fmt.Println()

	elapsed = time.Since(start)

	fmt.Printf("Zeroed map in %v, speed %vM/s\n", elapsed, float64(capacity)/elapsed.Seconds()/1000000)

	start = time.Now()

	var sum int64
	for i := 0; i < capacity; i++ {
		sum += int64(m[i])
		if i%trackFrequency == 0 {
			fmt.Print(".")
		}
	}
	fmt.Println()

	elapsed = time.Since(start)

	fmt.Printf("Summed in %v, speed %vM/s, sum is %v\n", elapsed, float64(capacity)/elapsed.Seconds()/1000000, sum)

	assert(err)
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	mmap "github.com/edsrzf/mmap-go"
)

type watch struct {
	start   time.Time
	elapsed time.Duration
}

func newWatch() watch {
	var w = watch{}
	w.Init()
	return w
}

func (w *watch) Init() {
	w.start = time.Now()
}

func (w *watch) Fix() {
	w.elapsed = time.Since(w.start)
}
func (w *watch) Restart() {
	w.Fix()
	w.Init()
}

func main() {
	var valueCount = 1 * 1024 * 1024
	var valueSize = 1024
	var capacity = valueCount * valueSize
	var transformCount = 1 * 1024 * 1024
	var trackFrequency = 128 * 1024

	var w = newWatch()

	f, err := os.OpenFile("large.data", os.O_RDWR|os.O_CREATE, 0600)
	assert(err)
	defer f.Close()

	f.WriteAt([]byte{0}, int64(capacity-1))

	w.Restart()
	fmt.Printf("File created in %v\n", w.elapsed)

	m, err := mmap.MapRegion(f, capacity, mmap.RDWR, 0, 0)
	assert(err)
	defer m.Unmap()

	w.Restart()
	fmt.Printf("Map created in %v\n", w.elapsed)

	var zeroes [1024]byte
	for i := 0; i < valueCount; i++ {
		copy(m[i*valueSize:(i+1)*valueSize], zeroes[:])
		if i%trackFrequency == 0 {
			fmt.Print(".")
			m.Flush()
		}
	}
	m.Flush()
	fmt.Println()

	w.Restart()
	fmt.Printf("Zeroed map in %v, speed %v MB/s\n", w.elapsed, float64(capacity)/w.elapsed.Seconds()/(1024*1024))

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var values [1024]byte
	for i := 0; i < transformCount; i++ {
		j := int(r.Int63() % int64(valueCount))
		copy(values[:], m[j*valueSize:(j+1)*valueSize])
		for k := 0; k < valueSize; k++ {
			values[k]++
		}
		copy(m[j*valueSize:(j+1)*valueSize], values[:])
		if i%trackFrequency == 0 {
			fmt.Print(".")
			m.Flush()
		}
	}
	m.Flush()
	fmt.Println()

	w.Restart()
	fmt.Printf("Incremented %v random bytes in %v, speed %v MB/s\n", transformCount, w.elapsed, float64(capacity)/w.elapsed.Seconds()/(1024*1024))

	var sum int64
	for i := 0; i < valueCount; i++ {
		copy(values[:], m[i*valueSize:(i+1)*valueSize])
		for k := 0; k < valueSize; k++ {
			sum += int64(values[k])
		}
		if i%trackFrequency == 0 {
			fmt.Print(".")
		}
	}
	fmt.Println()

	w.Restart()
	fmt.Printf("Summed in %v, speed %v MB/s, sum is %v\n", w.elapsed, float64(capacity)/w.elapsed.Seconds()/(1024*1024), sum)

	m.Flush()

	w.Restart()
	fmt.Printf("Flushed in %v\n", w.elapsed)
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

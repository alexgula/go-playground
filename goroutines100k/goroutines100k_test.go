package goroutines100k

import (
	"os"
	"testing"
)

var n = 100000

func TestMain(m *testing.M) {
	// fmt.Printf("GOMAXPROCS was %v\n", runtime.GOMAXPROCS(runtime.NumCPU()))
	// fmt.Printf("GOMAXPROCS now %v\n", runtime.GOMAXPROCS(-1))
	// debug.SetGCPercent(-1)
	os.Exit(m.Run())
}

func BenchmarkSeq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(runSeq(n))
	}
	b.ReportAllocs()
}

func BenchmarkParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(runParallel(n))
	}
	b.ReportAllocs()
}

func BenchmarkParallelFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(runParallelFast(n))
	}
	b.ReportAllocs()
}

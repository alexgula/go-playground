package factors

import (
	"math"
	"rand"
	"reflect"
	"testing"
	"testing/quick"
)

func TestAllOk(t *testing.T) {
}

func Test0(t *testing.T) {
	assertGenerate(t, 0, []int{})
}

func Test1(t *testing.T) {
	assertGenerate(t, 1, []int{})
}

func Test2(t *testing.T) {
	assertGenerate(t, 2, []int{2})
}

func Test3(t *testing.T) {
	assertGenerate(t, 3, []int{3})
}

func Test4(t *testing.T) {
	assertGenerate(t, 4, []int{2, 2})
}

func Test6(t *testing.T) {
	assertGenerate(t, 6, []int{2, 3})
}

func Test8(t *testing.T) {
	assertGenerate(t, 8, []int{2, 2, 2})
}

func Test18(t *testing.T) {
	assertGenerate(t, 18, []int{2, 3, 3})
}

func TestFactorsMultipleIsN(t *testing.T) {
	f := func(n int) bool {
		n = limit(abs(n), 1000000)
		factorList := generate(n)
		if n != mult(factorList) {
			t.Log(n)
			t.Log(mult(factorList))
			t.Log(factorList)
		}
		return n == mult(generate(n))
	}
	if err := quick.Check(f, &quick.Config{MaxCount: 10000, Rand: rand.New()}); err != nil {
		t.Error(err)
	}
}

func TestFactorsArePrimes(t *testing.T) {
	f := func(n int) bool {
		n = limit(abs(n), 1000000)
		for _, factor := range generate(n) {
			if !isPrime(factor) {
				return false
			}
		}
		return true
	}
	if err := quick.Check(f, &quick.Config{MaxCount: 10000}); err != nil {
		t.Error(err)
	}
}

func assertGenerate(t *testing.T, n int, expect []int) {
	var result = generate(n)
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Primefactors of %v should be %v (got %v)", n, expect, result)
	}
}

func mult(array []int) (multiple int) {
	multiple = 1
	for _, i := range array {
		multiple *= i
	}
	return
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func limit(n int, lim int) int {
	for n > lim {
		n /= 10
	}
	return n
}

func isPrime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

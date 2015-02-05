package factors

import (
	"math"
)

func generate(n int) (factorList []int) {
	factorList = []int{}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		for n%i == 0 {
			factorList = append(factorList, i)
			n /= i
		}
	}
	if n > 1 {
		factorList = append(factorList, n)
	}

	return
}

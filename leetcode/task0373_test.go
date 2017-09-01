package task0373

import (
	"math"
	"reflect"
	"testing"
)

func TestFindKSmallestPairs(t *testing.T) {
	cases := []struct {
		nums1  []int
		nums2  []int
		k      int
		expect [][]int
	}{
		{[]int{}, []int{}, 5, [][]int{}},
		{[]int{1, 7, 11}, []int{2, 4, 6}, 3, [][]int{{1, 2}, {1, 4}, {1, 6}}},
		{[]int{1, 4, 11}, []int{2, 4, 6}, 3, [][]int{{1, 2}, {1, 4}, {4, 2}}},
		{[]int{1, 4, 5}, []int{2, 4, 6}, 5, [][]int{{1, 2}, {1, 4}, {4, 2}, {5, 2}, {1, 6}}},
		{[]int{1, math.MaxInt32}, []int{1}, 5, [][]int{{1, 1}, {math.MaxInt32, 1}}},
		{[]int{-4, 3}, []int{2, 5}, 4, [][]int{{-4, 2}, {-4, 5}, {3, 2}, {3, 5}}},
		{[]int{1, 1, 2}, []int{1, 2, 3}, 10, [][]int{{1, 1}, {1, 1}, {2, 1}, {1, 2}, {1, 2}, {1, 3}, {2, 2}, {1, 3}, {2, 3}}},
	}

	for _, c := range cases {
		got := kSmallestPairs(c.nums1, c.nums2, c.k)
		if !reflect.DeepEqual(got, c.expect) {
			t.Errorf("nSmallestPairs(%v, %v, %v) == %v, expect %v", c.nums1, c.nums2, c.k, got, c.expect)
		}
	}
}

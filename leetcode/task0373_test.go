package task0373

import (
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
		{[]int{1, 7, 11}, []int{2, 4, 6}, 3, [][]int{{1, 2}, {1, 4}, {1, 6}}},
		{[]int{1, 4, 11}, []int{2, 4, 6}, 3, [][]int{{1, 2}, {1, 4}, {4, 2}}},
		{[]int{1, 4, 5}, []int{2, 4, 6}, 5, [][]int{{1, 2}, {1, 4}, {4, 2}, {1, 6}, {5, 2}}},
	}

	for _, c := range cases {
		got := findKSmallestPairs(c.nums1, c.nums2, c.k)
		if !reflect.DeepEqual(got, c.expect) {
			t.Errorf("nSmallestPairs(%v, %v, %v) == %v, expect %v", c.nums1, c.nums2, c.k, got, c.expect)
		}
	}
}

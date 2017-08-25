package task0373

import (
	"math"
)

func findKSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := make([][]int, k)

	var i1, i2 int

	for i := 0; i < len(result); i++ {
		result[i] = []int{nums1[i1], nums2[i2]}
		i1, i2 = findSmallestPairAfter(nums1, nums2, i1, i2)
	}

	return result
}

func findSmallestPairAfter(nums1, nums2 []int, p1, p2 int) (int, int) {
	var cur1, cur2 int
	var cursum = math.MaxInt32
	for i1 := 0; i1 < len(nums1); i1++ {
		var i2 int
		if i1 <= p1 {
			if p2 >= len(nums2)-1 {
				continue
			}
			i2 = p2 + 1
		} else {
			i2 = 0
		}
		var isum = pairSum(nums1, nums2, i1, i2)
		if isum < cursum {
			cur1, cur2, cursum = i1, i2, isum
		}
	}
	return cur1, cur2
}

func pairSum(nums1, nums2 []int, p1, p2 int) int {
	return nums1[p1] + nums2[p2]
}

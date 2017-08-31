package task0373

import "sort"

type bySum [][]int

func (a bySum) Len() int           { return len(a) }
func (a bySum) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySum) Less(i, j int) bool { return sum(a[i][0], a[i][1]) < sum(a[j][0], a[j][1]) }

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if k > len(nums1)*len(nums2) {
		k = len(nums1) * len(nums2)
	}

	result := make([][]int, k)

	var i, i1, i2 int

	for ; i < len(result); i++ {
		result[i] = []int{nums1[i1], nums2[i2]}
		i1, i2 = nextPair(nums1, nums2, i1, i2)
	}
	sort.Sort(bySum(result))

	for ; i < len(nums1)*len(nums2); i++ {
		insertPair(nums1, nums2, result, i1, i2)
		i1, i2 = nextPair(nums1, nums2, i1, i2)
	}
	sort.Sort(bySum(result))

	return result
}

func nextPair(nums1, nums2 []int, i1, i2 int) (int, int) {
	i1++
	if i1 >= len(nums1) {
		i1 = 0
		i2++
	}
	return i1, i2
}

func insertPair(nums1, nums2 []int, pairs [][]int, i1, i2 int) {
	var iMax, sumMax = findMax(pairs)
	if sumPair(nums1, nums2, i1, i2) < sumMax {
		pairs[iMax] = []int{nums1[i1], nums2[i2]}
	}
}

func findMax(pairs [][]int) (int, int64) {
	var iMax, sumMax = 0, sum(pairs[0][0], pairs[0][1])
	for i := 1; i < len(pairs); i++ {
		var sum = sum(pairs[i][0], pairs[i][1])
		if sum > sumMax {
			iMax, sumMax = i, sum
		}
	}
	return iMax, sumMax
}

func sumPair(nums1, nums2 []int, i1, i2 int) int64 {
	return sum(nums1[i1], nums2[i2])
}

func sum(val1, val2 int) int64 {
	return int64(val1) + int64(val2)
}

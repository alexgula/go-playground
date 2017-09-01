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

	if k == 0 {
		return result
	}

	var i, p1, p2 int

	for ; i < len(result); i++ {
		result[i] = []int{nums1[p1], nums2[p2]}
		p1, p2 = nextPair(nums1, nums2, p1, p2)
	}

	for j := len(result) / 2; j >= 0; j-- {
		heapify(bySum(result), j)
	}

	for ; i < len(nums1)*len(nums2); i++ {
		insertPair(nums1, nums2, result, p1, p2)
		p1, p2 = nextPair(nums1, nums2, p1, p2)
	}

	for j := len(result) / 2; j >= 0; j-- {
		var k = len(result) - j - 1
		result[j], result[k] = result[k], result[j]
	}
	sort.Sort(bySum(result))

	return result
}

func nextPair(nums1, nums2 []int, p1, p2 int) (int, int) {
	p1++
	if p1 >= len(nums1) {
		p1 = 0
		p2++
	}
	return p1, p2
}

func insertPair(nums1, nums2 []int, pairs [][]int, p1, p2 int) {
	var sumNew = sum(nums1[p1], nums2[p2])
	var sumCur = sum(pairs[0][0], pairs[0][1])
	if sumNew < sumCur {
		pairs[0][0], pairs[0][1] = nums1[p1], nums2[p2]
		heapify(bySum(pairs), 0)
	}
}

func heapify(data sort.Interface, i int) {
	var iLeft, iRight, iMax = 2*i + 1, 2*i + 2, i
	if iLeft < data.Len() && data.Less(iMax, iLeft) {
		iMax = iLeft
	}
	if iRight < data.Len() && data.Less(iMax, iRight) {
		iMax = iRight
	}
	if iMax != i {
		data.Swap(iMax, i)
		heapify(data, iMax)
	}
}

func sum(val1, val2 int) int64 {
	return int64(val1) + int64(val2)
}

package task0373

func findKSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := make([][]int, k)

	var i1, i2 int

	for i := 0; i < len(result); i++ {
		result[i] = []int{nums1[i1], nums2[i2]}
		if i1+1 == len(nums1) {
			i1 = 0
			i2++
		} else if i2+1 == len(nums2) {
			i1++
			i2 = 0
		} else if nums1[i1+1]+nums2[i2] < nums1[i1]+nums2[i2+1] {
			i1++
			i2 = 0
		} else {
			i1 = 0
			i2++
		}
	}

	return result
}

type pairs struct {
	nums1, nums2 []int
}

func (p pairs) getIth(i int) []int {
	return []int{p.nums1[i%len(p.nums1)], p.nums2[i/len(p.nums1)]}
}

func (p pairs) less(i1 int, i2 int) {
	return
}

func (p pairs) findKthSmallest(k, il, ir int) []int {
	if l
}

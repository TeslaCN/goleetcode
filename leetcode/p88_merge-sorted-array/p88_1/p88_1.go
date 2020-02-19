package p88_1

var Merge = merge

// Runtime: O(m+n)
// Space: O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m <= 0 {
		copy(nums1, nums2)
		return
	}
	cur, p1, p2 := len(nums1)-1, m-1, n-1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] < nums2[p2] {
			nums1[cur] = nums2[p2]
			p2--
		} else {
			nums1[cur] = nums1[p1]
			p1--
		}
		cur--
	}
	if p2 >= 0 {
		copy(nums1, nums2[:p2+1])
	}
}

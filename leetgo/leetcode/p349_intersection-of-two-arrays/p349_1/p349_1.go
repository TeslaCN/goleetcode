package p349_1

import "sort"

// Runtime: O(nlog(n))
// Space: O(1)
func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	ans := make([]int, 0)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if value := nums1[i]; nums1[i] == nums2[j] {
			ans = append(ans, value)
			for ; i < len(nums1) && nums1[i] == value; i++ {
			}
			for ; j < len(nums2) && nums2[j] == value; j++ {
			}
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 89.38% 的用户
内存消耗 : 2.8 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/

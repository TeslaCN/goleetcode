package p215_3

var FindKthLargest = findKthLargest

// heap
// Runtime: O(nlog(n))
// Space: O(1)
func findKthLargest(nums []int, k int) int {
	for i := len(nums)>>1 - 1; i >= 0; i-- {
		adjustHeap(nums, i, len(nums))
	}
	for i := len(nums) - 1; i > 0; i-- {
		k--
		if k == 0 {
			return nums[0]
		}
		nums[0], nums[i] = nums[i], nums[0]
		adjustHeap(nums, 0, i)
	}
	return nums[0]
}

func adjustHeap(heap []int, pos, length int) {
	element := heap[pos]
	for child := pos<<1 + 1; child < length; child = child<<1 + 1 {
		if child+1 < length && heap[child+1] > heap[child] {
			child++
		}
		if element < heap[child] {
			heap[pos] = heap[child]
			pos = child
		}
	}
	heap[pos] = element
}

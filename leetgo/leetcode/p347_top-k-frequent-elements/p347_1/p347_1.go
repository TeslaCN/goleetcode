package p347_1

// heap sort
// Runtime: O(n*log(n))
// Space: O(n)
func topKFrequent(nums []int, k int) []int {
	if k <= 0 || len(nums) <= 0 {
		return []int{}
	}
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	if k > 1 {
		heap := make([]int, 0)
		res := make([]int, 0, k)
		for key := range m {
			heap = append(heap, key)
		}
		for i := len(heap)>>1 - 1; i >= 0; i-- {
			adjustHeap(heap, i, len(heap), m)
		}
		for i := len(heap) - 1; i > len(heap)-1-k; i-- {
			res = append(res, heap[0])
			heap[0], heap[i] = heap[i], heap[0]
			adjustHeap(heap, 0, i, m)
		}
		return res
	} else {
		maxPos, maxVal := nums[0], m[nums[0]]
		for pos, val := range m {
			if val > maxVal {
				maxVal = val
				maxPos = pos
			}
		}
		return []int{maxPos}
	}
}

func adjustHeap(heap []int, pos, length int, priority map[int]int) {
	num := heap[pos]
	for child := pos<<1 + 1; child < length; child = child<<1 + 1 {
		if child < length-1 && priority[heap[child]] < priority[heap[child+1]] {
			child++
		}
		if priority[num] < priority[heap[child]] {
			heap[pos] = heap[child]
			pos = child
		} else {
			break
		}
	}
	heap[pos] = num
}

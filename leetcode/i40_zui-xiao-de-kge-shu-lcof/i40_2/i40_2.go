package i40_2

var GetLeastNumbers = getLeastNumbers

func getLeastNumbers(arr []int, k int) []int {
	n := len(arr)
	for i := (n - 1) >> 1; i >= 0; i-- {
		adjustHeap(arr, i, n)
	}
	for i := n - 1; i >= n-k; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		adjustHeap(arr, 0, i)
	}
	return arr[n-k:]
}

func adjustHeap(arr []int, target, end int) {
	value := arr[target]
	for i := target<<1 + 1; i < end; i = i<<1 + 1 {
		if i+1 < end && arr[i] > arr[i+1] {
			i++
		}
		if arr[i] < value {
			arr[target] = arr[i]
			target = i
		}
	}
	arr[target] = value
}

/*
       0
     1    2
   3  4  5  6
*/

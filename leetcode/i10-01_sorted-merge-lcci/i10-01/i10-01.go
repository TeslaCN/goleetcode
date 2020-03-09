package i10_01

func merge(A []int, m int, B []int, n int) {
	i := len(A) - 1
	for m > 0 && n > 0 {
		if A[m-1] > B[n-1] {
			A[i] = A[m-1]
			m--
		} else {
			A[i] = B[n-1]
			n--
		}
		i--
	}
	for n > 0 {
		A[i] = B[n-1]
		i--
		n--
	}
}

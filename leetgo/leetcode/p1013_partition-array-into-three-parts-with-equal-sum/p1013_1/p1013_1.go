package p1013_1

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, n := range A {
		sum += n
	}
	if sum%3 != 0 {
		return false
	}
	target := sum / 3
	sum = 0
	cur := 0
	var pos [2]int

	for i, value := range A {
		sum += value
		if sum == target && cur < 2 {
			pos[cur] = i
			cur++
			sum = 0
		}
	}
	return sum == target && pos[0] < pos[1] && pos[1] < len(A)-1
}

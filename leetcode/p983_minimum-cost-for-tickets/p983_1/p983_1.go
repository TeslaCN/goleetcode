package p983_1

// Deprecated:TLE
var MincostTickets = mincostTickets

var exp = [3]int{1, 7, 30}

// Deprecated
func mincostTickets(days []int, costs []int) int {
	switch len(days) {
	case 0:
		return 0
	case 1:
		return costs[0]
	}
	minValue := -1
	for i := 0; i < 3; i++ {
		maxDay := days[0] + exp[i]
		next := 0
		for next < len(days) && days[next] < maxDay {
			next++
		}
		if cost := costs[i] + mincostTickets(days[next:], costs); minValue == -1 || cost < minValue {
			minValue = cost
		}
	}
	return minValue
}

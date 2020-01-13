package p121_1

//
// Deprecated
func MaxProfit(prices []int) int {
	if prices == nil {
		return 0
	}
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			value := prices[j] - prices[i]
			if value > max {
				max = value
			}
		}
	}
	return max
}

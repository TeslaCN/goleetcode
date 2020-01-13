package p121_2

func MaxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}
	minValue := prices[0]
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minValue {
			minValue = prices[i]
		} else if profit := prices[i] - minValue; profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
